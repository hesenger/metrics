# Tasks: Add Frontend Authentication

## Backend Changes

- [x] 1. Update JWT expiration from 24 hours to 7 days in token generation
- [x] 2. Update login handler to set JWT in HttpOnly cookie named "token" with Strict SameSite, 7-day expiration
- [x] 3. Update login handler response to exclude token from body, return only user object
- [x] 4. Update register handler to set JWT in HttpOnly cookie with same flags as login
- [x] 5. Update register handler response to exclude token from body, return only user object
- [x] 6. Update Google OAuth callback to set JWT in HttpOnly cookie
- [x] 7. Update Google OAuth callback to redirect to frontend with user data in query params (id, email, oauth_provider)
- [x] 8. Update Google OAuth callback error handling to redirect to `/login?error=...`
- [x] 9. Add GET `/api/auth/me` endpoint that validates JWT from cookie
- [x] 10. Implement /me endpoint to return user data (id, email, oauth_provider, created_at)
- [x] 11. Implement /me endpoint to refresh JWT cookie expiration on each request (extends session)
- [x] 12. Implement /me endpoint to return 401 if JWT is invalid or expired
- [x] 13. Add logout endpoint POST `/api/auth/logout` that clears the token cookie
- [x] 14. Test all auth endpoints verify cookie is set with correct flags (HttpOnly, SameSite=Strict, 7-day expiration)
- [x] 15. Test /me endpoint validates JWT and refreshes cookie
- [x] 16. Test /me endpoint returns 401 for missing or invalid JWT
- [x] 17. Test logout endpoint clears cookie properly

## Frontend Setup

- [x] 18. Install @tanstack/react-query dependency via bun
- [x] 19. Create QueryClient and wrap app with QueryClientProvider in main.tsx
- [x] 20. Create types/user.ts with User interface (id, email, oauth_provider, created_at)
- [x] 21. Create contexts/auth-context.tsx with AuthContext, AuthProvider, and useAuth hook
- [x] 22. Implement auth context to load initial state from localStorage on mount
- [x] 23. Implement auth context to call /me endpoint on mount if localStorage has user data
- [x] 24. Handle /me response in auth context: update localStorage and state on success
- [x] 25. Handle /me error in auth context: clear localStorage and set unauthenticated on 401
- [x] 26. Update main.tsx to wrap app with AuthProvider (inside MantineProvider)
- [x] 27. Create api/auth.ts with mutation and query functions
- [x] 28. Create useLoginMutation hook in api/auth.ts
- [x] 29. Create useRegisterMutation hook in api/auth.ts
- [x] 30. Create useLogoutMutation hook in api/auth.ts
- [x] 31. Create useMeQuery hook in api/auth.ts for GET /api/auth/me

## Protected Routes

- [x] 32. Create components/protected-route.tsx that checks isAuthenticated from useAuth
- [x] 33. Update App.tsx to wrap home route with ProtectedRoute component
- [x] 34. Test unauthenticated users are redirected to /login when accessing home

## Login Page

- [x] 35. Create pages/login.tsx with login form layout
- [x] 36. Add email and password input fields using Mantine components
- [x] 37. Add form validation for email format and required fields
- [x] 38. Add submit button that triggers useLoginMutation
- [x] 39. Display loading state during mutation
- [x] 40. Display error messages from API or validation
- [x] 41. Add "Sign in with Google" button that redirects to /api/auth/google
- [x] 42. Add link to /register page
- [x] 43. Add redirect logic for authenticated users to home page
- [x] 44. Update App.tsx to add /login route
- [x] 45. Test login flow with valid credentials stores user in localStorage and redirects
- [x] 46. Test login with invalid credentials shows error message

## Registration Page

- [x] 47. Create pages/register.tsx with registration form layout
- [x] 48. Add email and password input fields using Mantine components
- [x] 49. Add form validation for email format, password length (min 8), required fields
- [x] 50. Add submit button that triggers useRegisterMutation
- [x] 51. Display loading state during mutation
- [x] 52. Display error messages from API or validation
- [x] 53. Add "Sign in with Google" button that redirects to /api/auth/google
- [x] 54. Add link to /login page
- [x] 55. Add redirect logic for authenticated users to home page
- [x] 56. Update App.tsx to add /register route
- [x] 57. Test registration flow with valid data creates user and redirects
- [x] 58. Test registration with existing email shows conflict error
- [x] 59. Test registration with weak password shows validation error

## Google OAuth Callback

- [x] 60. Create pages/auth-callback.tsx to handle OAuth redirect from backend
- [x] 61. Extract user data from URL query parameters (id, email, oauth_provider)
- [x] 62. Store user data in localStorage under "user" key
- [x] 63. Update auth context with user data
- [x] 64. Redirect to home page after successful data extraction
- [x] 65. Handle missing query params by showing error and redirecting to /login
- [x] 66. Update App.tsx to add /auth/callback route (public)
- [x] 67. Test Google OAuth flow end-to-end with backend callback

## Logout Functionality

- [x] 68. Add logout button to home page or navigation
- [x] 69. Wire logout button to useLogoutMutation
- [x] 70. Ensure logout mutation clears localStorage "user" key
- [x] 71. Ensure logout mutation updates auth context to null
- [x] 72. Ensure logout redirects to /login page
- [x] 73. Test logout clears all auth state and cookie

## Integration Testing

- [x] 74. Test full registration flow: register -> authenticated -> home page
- [x] 75. Test full login flow: login -> authenticated -> home page
- [x] 76. Test full Google OAuth flow: click button -> OAuth -> callback -> home page
- [x] 77. Test logout flow: logout -> unauthenticated -> redirected to login
- [x] 78. Test protected route: access home without auth -> redirected to login
- [x] 79. Test /me endpoint called on app mount and refreshes auth state
- [x] 80. Test /me endpoint failure clears auth state and redirects to login
- [x] 81. Test localStorage persistence: refresh page while authenticated -> /me validates -> still authenticated
- [x] 82. Test auth pages redirect: access /login while authenticated -> redirected to home
- [x] 83. Test session extension: use app within 7 days -> session extended via /me calls
- [x] 84. Test session expiration: wait 7 days without activity -> /me returns 401 -> redirected to login

## Documentation

- [x] 85. Update README or docs with new authentication flow (if README exists)
- [x] 86. Verify all environment variables are documented (JWT_SECRET, GOOGLE_*)
- [x] 87. Document /me endpoint behavior and 7-day session expiration
