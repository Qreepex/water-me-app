package types

import "github.com/golang-jwt/jwt/v5"

// ValidationError mirrors the TS validation shape.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// User represents a user account.
type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	Language     string `json:"language"`
	CreatedAt    string `json:"createdAt"`
}

// SignupRequest is the request body for user registration.
type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest is the request body for user authentication.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse contains the JWT token after successful login.
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// UpdateUserRequest is for updating user profile information.
type UpdateUserRequest struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Language *string `json:"language,omitempty"`
}

// ChangePasswordRequest is for changing user password.
type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type Claims struct {
	UserID string `json:"uid"`
	jwt.RegisteredClaims
}
