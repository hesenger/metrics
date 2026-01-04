package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	backend "github.com/hesen/metrics"
	"github.com/hesen/metrics/internal/config"
	"github.com/hesen/metrics/internal/database"
	"github.com/hesen/metrics/internal/handlers"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := database.RunMigrations(cfg.DatabaseURL); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	queries := database.New(pool)
	authHandler := handlers.NewAuthHandler(queries, cfg.JWTSecret, cfg.GoogleClientID, cfg.GoogleClientSecret, cfg.GoogleRedirectURL)

	app := fiber.New()

	app.Post("/api/auth/register", authHandler.Register)
	app.Post("/api/auth/login", authHandler.Login)
	app.Get("/api/auth/google", authHandler.InitiateGoogleOAuth)
	app.Get("/api/auth/google/callback", authHandler.GoogleOAuthCallback)
	app.Get("/api/auth/me", authHandler.Me)
	app.Post("/api/auth/logout", authHandler.Logout)

	webFS, err := fs.Sub(backend.WebAssets, "web")
	if err != nil {
		log.Fatalf("failed to create web filesystem: %v", err)
	}

	app.Use("*", filesystem.New(filesystem.Config{
		Root:  http.FS(webFS),
		Index: "index.html",
	}))

	log.Printf("server starting on port %s", cfg.Port)
	if err := app.Listen(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
