# frontend-authentication Specification

## Purpose
TBD - to be updated after archiving

## ADDED Requirements

### Requirement: Login page
The system MUST provide a login page for users to authenticate with email/password or Google OAuth.

#### Scenario: Login page accessible at /login
**Given** the frontend application
**Then** a route exists at `/login`
**And** the page renders a login form
**And** the page is accessible to unauthenticated users
**And** authenticated users are redirected to home page

#### Scenario: Email/password login form
**Given** a user on the `/login` page
**Then** the page displays an email input field
**And** the page displays a password input field
**And** the page displays a "Login" submit button
**And** the page displays a "Sign in with Google" button
**And** the page displays a link to `/register` page

#### Scenario: Successful email/password login
**Given** a user enters valid credentials on `/login` page
**When** the user submits the login form
**Then** a POST request is made to `/api/auth/login` via React Query mutation
**And** the JWT cookie is set by the backend
**And** the user data is stored in localStorage under key "user"
**And** the auth context updates with authenticated state
**And** the user is redirected to the home page

#### Scenario: Failed login shows error
**Given** a user enters invalid credentials
**When** the user submits the login form
**And** the API returns 401 Unauthorized
**Then** an error message is displayed on the page
**And** the user remains on the login page
**And** no data is stored in localStorage

#### Scenario: Google OAuth login initiation
**Given** a user on the `/login` page
**When** the user clicks "Sign in with Google" button
**Then** the browser redirects to `/api/auth/google`
**And** the backend redirects to Google's OAuth consent screen

#### Scenario: Form validation on login
**Given** a user on the `/login` page
**When** the user submits the form with empty fields
**Then** validation errors are displayed inline
**And** no API request is made until validation passes

### Requirement: Registration page
The system MUST provide a registration page for new users to create accounts.

#### Scenario: Registration page accessible at /register
**Given** the frontend application
**Then** a route exists at `/register`
**And** the page renders a registration form
**And** the page is accessible to unauthenticated users
**And** authenticated users are redirected to home page

#### Scenario: Email/password registration form
**Given** a user on the `/register` page
**Then** the page displays an email input field
**And** the page displays a password input field
**And** the page displays a "Register" submit button
**And** the page displays a "Sign in with Google" button
**And** the page displays a link to `/login` page

#### Scenario: Successful registration
**Given** a user enters valid registration details
**When** the user submits the registration form
**Then** a POST request is made to `/api/auth/register` via React Query mutation
**And** the JWT cookie is set by the backend
**And** the user data is stored in localStorage under key "user"
**And** the auth context updates with authenticated state
**And** the user is redirected to the home page

#### Scenario: Failed registration shows error
**Given** a user attempts to register with an existing email
**When** the user submits the form
**And** the API returns 409 Conflict
**Then** an error message is displayed indicating email is taken
**And** the user remains on the registration page
**And** no data is stored in localStorage

#### Scenario: Form validation on registration
**Given** a user on the `/register` page
**When** the user enters a password less than 8 characters
**Then** a validation error is displayed
**And** the submit button is disabled or shows error on submit
**When** the user enters an invalid email format
**Then** a validation error is displayed for email field

### Requirement: Authentication context
The system MUST provide a global authentication context for managing auth state across the application.

#### Scenario: Auth context wraps entire app
**Given** the frontend application
**Then** an AuthProvider component wraps the app in `main.tsx`
**And** all components can access auth context via `useAuth` hook
**And** the context provides `user` object or null
**And** the context provides `isAuthenticated` boolean

#### Scenario: Auth context loads from localStorage on mount
**Given** the application loads
**And** localStorage contains user data under key "user"
**Then** the auth context initializes with authenticated state
**And** the user object is populated from localStorage
**When** localStorage is empty
**Then** the auth context initializes as unauthenticated

#### Scenario: Auth context updates on login
**Given** a user successfully logs in or registers
**When** the mutation succeeds
**Then** the auth context `setUser` function is called with user data
**And** `isAuthenticated` becomes true
**And** the user object contains id, email, and oauth_provider fields

#### Scenario: Auth context clears on logout
**Given** a user is authenticated
**When** logout is triggered
**Then** the auth context `setUser` is called with null
**And** `isAuthenticated` becomes false
**And** localStorage "user" key is removed

### Requirement: Protected routes
The system MUST restrict access to certain routes for authenticated users only.

#### Scenario: Protected route wrapper component
**Given** the application routing configuration
**Then** a ProtectedRoute component exists
**And** the component checks `isAuthenticated` from auth context
**And** the component accepts a child component to render

#### Scenario: Unauthenticated access redirects to login
**Given** a user is not authenticated
**When** the user attempts to access a protected route
**Then** the ProtectedRoute component redirects to `/login`
**And** the protected component does not render

#### Scenario: Authenticated access allows through
**Given** a user is authenticated
**When** the user accesses a protected route
**Then** the ProtectedRoute component renders the child component
**And** no redirect occurs

#### Scenario: Home page is protected
**Given** the route configuration
**Then** the `/` route is wrapped with ProtectedRoute
**And** unauthenticated users are redirected to `/login`

### Requirement: Logout functionality
The system MUST allow authenticated users to log out and clear their session.

#### Scenario: Logout action clears auth state
**Given** a user is authenticated
**When** the user triggers logout
**Then** a POST request is made to `/api/auth/logout`
**And** localStorage "user" key is removed
**And** the auth context is cleared
**And** the user is redirected to `/login` page

#### Scenario: Logout button accessible when authenticated
**Given** a user is authenticated
**Then** a logout button or link is visible in the UI
**When** the user clicks logout
**Then** the logout flow is triggered

### Requirement: React Query integration
The system MUST use React Query for auth-related API calls and state management.

#### Scenario: React Query provider wraps app
**Given** the frontend application
**Then** QueryClientProvider wraps the app
**And** a QueryClient instance is created
**And** all components can use React Query hooks

#### Scenario: Login mutation using React Query
**Given** the auth API module
**Then** a `useLoginMutation` hook exists
**And** the hook uses `useMutation` from React Query
**And** the mutation POSTs to `/api/auth/login`
**And** the mutation handles success and error states

#### Scenario: Register mutation using React Query
**Given** the auth API module
**Then** a `useRegisterMutation` hook exists
**And** the hook uses `useMutation` from React Query
**And** the mutation POSTs to `/api/auth/register`
**And** the mutation handles success and error states

#### Scenario: Logout mutation using React Query
**Given** the auth API module
**Then** a `useLogoutMutation` hook exists
**And** the hook uses `useMutation` from React Query
**And** the mutation POSTs to `/api/auth/logout`
**And** the mutation handles success callback to clear state

### Requirement: Google OAuth callback handling
The system MUST handle the OAuth callback from Google and extract user data from URL parameters.

#### Scenario: Callback URL contains user data
**Given** a user completes Google OAuth
**When** the backend redirects to frontend callback URL
**Then** the URL contains user data as query parameters (id, email, oauth_provider)
**And** the frontend extracts user data from URL
**And** user data is stored in localStorage
**And** auth context is updated with authenticated state
**And** user is redirected to home page

#### Scenario: Callback URL missing user data shows error
**Given** the OAuth callback URL is accessed
**When** required query parameters are missing
**Then** an error message is displayed
**And** the user is redirected to `/login` page
**And** no auth state is set
