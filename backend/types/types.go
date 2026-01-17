package types

import "github.com/golang-jwt/jwt/v5"

// Plant mirrors the frontend Plant type.
type Plant struct {
	ID                      string              `json:"id"`
	UserID                  string              `json:"userId"`
	Species                 string              `json:"species"`
	Name                    string              `json:"name"`
	SunLight                SunlightRequirement `json:"sunLight"`
	PreferedTemperature     float64             `json:"preferedTemperature"`
	WateringIntervalDays    int                 `json:"wateringIntervalDays"`
	LastWatered             string              `json:"lastWatered"`
	FertilizingIntervalDays int                 `json:"fertilizingIntervalDays"`
	LastFertilized          string              `json:"lastFertilized"`
	PreferedHumidity        float64             `json:"preferedHumidity"`
	SprayIntervalDays       *int                `json:"sprayIntervalDays,omitempty"`
	Notes                   []string            `json:"notes"`
	Flags                   []PlantFlag         `json:"flags"`
	PhotoIDs                []string            `json:"photoIds"`
}

// PlantInput keeps pointers so we can detect missing fields for PATCH.
type PlantInput struct {
	ID                      *string              `json:"id,omitempty"`
	Species                 *string              `json:"species"`
	Name                    *string              `json:"name"`
	SunLight                *SunlightRequirement `json:"sunLight"`
	PreferedTemperature     *float64             `json:"preferedTemperature"`
	WateringIntervalDays    *int                 `json:"wateringIntervalDays"`
	LastWatered             *string              `json:"lastWatered"`
	FertilizingIntervalDays *int                 `json:"fertilizingIntervalDays"`
	LastFertilized          *string              `json:"lastFertilized"`
	PreferedHumidity        *float64             `json:"preferedHumidity"`
	SprayIntervalDays       *int                 `json:"sprayIntervalDays"`
	Notes                   *[]string            `json:"notes"`
	Flags                   *[]PlantFlag         `json:"flags"`
	PhotoIDs                *[]string            `json:"photoIds"`
}

type PlantFlag string

const (
	PlantFlagNoDraught         PlantFlag = "No Draught"
	PlantFlagRemoveBrownLeaves PlantFlag = "Remove Brown Leaves"
)

type SunlightRequirement string

const (
	SunlightFullSun            SunlightRequirement = "Full Sun"
	SunlightIndirectSun        SunlightRequirement = "Indirect Sun"
	SunlightPartialShade       SunlightRequirement = "Partial Shade"
	SunlightPartialToFullShade SunlightRequirement = "Partial to Full Shade"
	SunlightFullShade          SunlightRequirement = "Full Shade"
)

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
