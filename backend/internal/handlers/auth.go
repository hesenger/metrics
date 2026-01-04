package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/hesen/metrics/internal/auth"
	"github.com/hesen/metrics/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	oauth2api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type AuthHandler struct {
	queries      *database.Queries
	jwtSecret    string
	googleOAuth  *oauth2.Config
	stateStore   *auth.StateStore
}

func NewAuthHandler(queries *database.Queries, jwtSecret string, googleClientID, googleClientSecret, googleRedirectURL string) *AuthHandler {
	return &AuthHandler{
		queries:   queries,
		jwtSecret: jwtSecret,
		googleOAuth: &oauth2.Config{
			ClientID:     googleClientID,
			ClientSecret: googleClientSecret,
			RedirectURL:  googleRedirectURL,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
		stateStore: auth.NewStateStore(),
	}
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	User UserDetails `json:"user"`
}

type UserDetails struct {
	ID            int64  `json:"id"`
	Email         string `json:"email"`
	OAuthProvider string `json:"oauth_provider,omitempty"`
	CreatedAt     string `json:"created_at"`
}

func (h *AuthHandler) setJWTCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Strict",
		MaxAge:   7 * 24 * 60 * 60,
		Path:     "/",
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and password are required",
		})
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid email format",
		})
	}

	if len(req.Password) < 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password must be at least 8 characters",
		})
	}

	passwordHash, err := auth.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	user, err := h.queries.CreateUser(context.Background(), database.CreateUserParams{
		Email:        req.Email,
		PasswordHash: pgtype.Text{String: passwordHash, Valid: true},
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "email already registered",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	token, err := auth.GenerateToken(user.ID, user.Email, h.jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	h.setJWTCookie(c, token)

	oauthProvider := ""
	if user.OauthProvider.Valid {
		oauthProvider = user.OauthProvider.String
	}

	return c.Status(fiber.StatusCreated).JSON(AuthResponse{
		User: UserDetails{
			ID:            user.ID,
			Email:         user.Email,
			OAuthProvider: oauthProvider,
			CreatedAt:     user.CreatedAt.Time.Format("2006-01-02T15:04:05Z"),
		},
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "email and password are required",
		})
	}

	user, err := h.queries.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid credentials",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}

	if err := auth.VerifyPassword(user.PasswordHash.String, req.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	token, err := auth.GenerateToken(user.ID, user.Email, h.jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	h.setJWTCookie(c, token)

	oauthProvider := ""
	if user.OauthProvider.Valid {
		oauthProvider = user.OauthProvider.String
	}

	return c.Status(fiber.StatusOK).JSON(AuthResponse{
		User: UserDetails{
			ID:            user.ID,
			Email:         user.Email,
			OAuthProvider: oauthProvider,
			CreatedAt:     user.CreatedAt.Time.Format("2006-01-02T15:04:05Z"),
		},
	})
}

func (h *AuthHandler) InitiateGoogleOAuth(c *fiber.Ctx) error {
	state, err := h.stateStore.GenerateState()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to generate state",
		})
	}

	url := h.googleOAuth.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func (h *AuthHandler) GoogleOAuthCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if !h.stateStore.ValidateState(state) {
		return c.Redirect("/login?error=invalid_state", fiber.StatusFound)
	}

	code := c.Query("code")
	if code == "" {
		return c.Redirect("/login?error=oauth_failed", fiber.StatusFound)
	}

	token, err := h.googleOAuth.Exchange(context.Background(), code)
	if err != nil {
		return c.Redirect("/login?error=oauth_failed", fiber.StatusFound)
	}

	oauth2Service, err := oauth2api.NewService(context.Background(), option.WithTokenSource(h.googleOAuth.TokenSource(context.Background(), token)))
	if err != nil {
		return c.Redirect("/login?error=oauth_failed", fiber.StatusFound)
	}

	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return c.Redirect("/login?error=oauth_failed", fiber.StatusFound)
	}

	existingUser, err := h.queries.GetUserByOAuthProvider(context.Background(), database.GetUserByOAuthProviderParams{
		OauthProvider: pgtype.Text{String: "google", Valid: true},
		OauthID:       pgtype.Text{String: userInfo.Id, Valid: true},
	})

	var userID int64
	var userEmail string

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			newUser, err := h.queries.CreateOAuthUser(context.Background(), database.CreateOAuthUserParams{
				Email:         userInfo.Email,
				OauthProvider: pgtype.Text{String: "google", Valid: true},
				OauthID:       pgtype.Text{String: userInfo.Id, Valid: true},
			})
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "failed to create user",
				})
			}
			userID = newUser.ID
			userEmail = newUser.Email
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to fetch user",
			})
		}
	} else {
		userID = existingUser.ID
		userEmail = existingUser.Email
	}

	jwtToken, err := auth.GenerateToken(userID, userEmail, h.jwtSecret)
	if err != nil {
		return c.Redirect("/login?error=oauth_failed", fiber.StatusFound)
	}

	h.setJWTCookie(c, jwtToken)

	redirectURL := fmt.Sprintf("/auth/callback?id=%d&email=%s&oauth_provider=google", userID, url.QueryEscape(userEmail))
	return c.Redirect(redirectURL, fiber.StatusFound)
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing authentication",
		})
	}

	claims, err := auth.ValidateToken(tokenString, h.jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid or expired token",
		})
	}

	user, err := h.queries.GetUserByID(context.Background(), claims.UserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "user not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch user",
		})
	}

	newToken, err := auth.GenerateToken(user.ID, user.Email, h.jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to refresh token",
		})
	}

	h.setJWTCookie(c, newToken)

	oauthProvider := ""
	if user.OauthProvider.Valid {
		oauthProvider = user.OauthProvider.String
	}

	return c.Status(fiber.StatusOK).JSON(AuthResponse{
		User: UserDetails{
			ID:            user.ID,
			Email:         user.Email,
			OAuthProvider: oauthProvider,
			CreatedAt:     user.CreatedAt.Time.Format("2006-01-02T15:04:05Z"),
		},
	})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		HTTPOnly: true,
		Secure:   c.Protocol() == "https",
		SameSite: "Strict",
		MaxAge:   -1,
		Path:     "/",
	})

	return c.SendStatus(fiber.StatusNoContent)
}
