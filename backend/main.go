package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ctx := context.Background()
	connString := getenv("DATABASE_URL", "postgres://postgres:pw@localhost:5432/plants?sslmode=disable")
	jwtSecret := getenv("JWT_SECRET", "your-secret-key-change-this-in-production")

	store, err := NewStore(ctx, connString)
	if err != nil {
		log.Fatalf("failed to initialize store: %v", err)
	}
	defer store.Close()

	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/api/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != http.MethodPost {
			methodNotAllowed(w)
			return
		}
		handleSignup(store, jwtSecret, w, r)
	})

	mux.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method != http.MethodPost {
			methodNotAllowed(w)
			return
		}
		handleLogin(store, jwtSecret, w, r)
	})

	// Protected user routes
	mux.HandleFunc("/api/user", authMiddleware(jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		if handlePreflight(w, r, []string{http.MethodGet, http.MethodPut}) {
			return
		}
		switch r.Method {
		case http.MethodGet:
			handleGetUser(store, w, r)
		case http.MethodPut:
			handleUpdateUserProfile(store, w, r)
		default:
			methodNotAllowed(w)
		}
	}))

	mux.HandleFunc("/api/user/password", authMiddleware(jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		if handlePreflight(w, r, []string{http.MethodPut}) {
			return
		}
		switch r.Method {
		case http.MethodPut:
			handleChangePassword(store, w, r)
		default:
			methodNotAllowed(w)
		}
	}))

	// Protected plant routes
	mux.HandleFunc("/api/plants", authMiddleware(jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		if handlePreflight(w, r, []string{http.MethodGet, http.MethodPost}) {
			return
		}
		switch r.Method {
		case http.MethodGet:
			handleListPlants(store, w, r)
		case http.MethodPost:
			handleCreatePlant(store, w, r)
		default:
			methodNotAllowed(w)
		}
	}))

	mux.HandleFunc("/api/plants/", authMiddleware(jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/plants/")
		if id == "" || strings.Contains(id, "/") {
			notFound(w)
			return
		}

		if handlePreflight(w, r, []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodDelete}) {
			return
		}

		switch r.Method {
		case http.MethodGet:
			handleGetPlant(store, id, w, r)
		case http.MethodPut:
			handleUpdatePlant(store, id, false, w, r)
		case http.MethodPatch:
			handleUpdatePlant(store, id, true, w, r)
		case http.MethodDelete:
			handleDeletePlant(store, id, w, r)
		default:
			methodNotAllowed(w)
		}
	}))

	addr := getenv("PORT", "8080")
	srv := &http.Server{
		Addr:         ":" + addr,
		Handler:      corsMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func handleListPlants(store *Store, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	plants, err := store.ListPlants(r.Context(), userID)
	if err != nil {
		serverError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, plants)
}

func handleGetPlant(store *Store, id string, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	plant, found, err := store.GetPlant(r.Context(), id, userID)
	if err != nil {
		serverError(w, err)
		return
	}
	if !found {
		notFound(w)
		return
	}
	respondJSON(w, http.StatusOK, plant)
}

func handleCreatePlant(store *Store, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	var input PlantInput
	if err := decodeJSON(r, &input); err != nil {
		badRequest(w, err.Error(), nil)
		return
	}

	sanitized := SanitizePlantInput(input)
	errors := ValidatePlantInput(sanitized, false)
	if len(errors) > 0 {
		badRequest(w, "Validation failed", errors)
		return
	}

	plant, err := store.CreatePlant(r.Context(), userID, sanitized)
	if err != nil {
		serverError(w, err)
		return
	}
	respondJSON(w, http.StatusCreated, plant)
}

func handleUpdatePlant(store *Store, id string, partial bool, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	var input PlantInput
	if err := decodeJSON(r, &input); err != nil {
		badRequest(w, err.Error(), nil)
		return
	}

	sanitized := SanitizePlantInput(input)
	errors := ValidatePlantInput(sanitized, partial)
	if len(errors) > 0 {
		badRequest(w, "Validation failed", errors)
		return
	}

	plant, found, err := store.UpdatePlant(r.Context(), id, userID, sanitized)
	if err != nil {
		serverError(w, err)
		return
	}
	if !found {
		notFound(w)
		return
	}
	respondJSON(w, http.StatusOK, plant)
}

func handleDeletePlant(store *Store, id string, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	deleted, err := store.DeletePlant(r.Context(), id, userID)
	if err != nil {
		serverError(w, err)
		return
	}
	if !deleted {
		notFound(w)
		return
	}
	respondJSON(w, http.StatusOK, map[string]bool{"success": true})
}

func decodeJSON(r *http.Request, dest any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return nil
}

func handleSignup(store *Store, jwtSecret string, w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := decodeJSON(r, &req); err != nil {
		badRequest(w, err.Error(), nil)
		return
	}

	if req.Email == "" || req.Password == "" {
		badRequest(w, "Email and password are required", nil)
		return
	}

	if len(req.Password) < 6 {
		badRequest(w, "Password must be at least 6 characters", nil)
		return
	}

	// Check if user already exists
	_, exists, err := store.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		serverError(w, err)
		return
	}
	if exists {
		badRequest(w, "User with this email already exists", nil)
		return
	}

	// Hash password
	passwordHash, err := HashPassword(req.Password)
	if err != nil {
		serverError(w, err)
		return
	}

	// Create user
	user, err := store.CreateUser(r.Context(), req.Email, passwordHash)
	if err != nil {
		serverError(w, err)
		return
	}

	// Generate JWT
	token, err := GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		serverError(w, err)
		return
	}

	respondJSON(w, http.StatusCreated, LoginResponse{
		Token: token,
		User:  user,
	})
}

func handleLogin(store *Store, jwtSecret string, w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := decodeJSON(r, &req); err != nil {
		badRequest(w, err.Error(), nil)
		return
	}

	if req.Email == "" || req.Password == "" {
		badRequest(w, "Email and password are required", nil)
		return
	}

	// Get user by email
	user, exists, err := store.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		serverError(w, err)
		return
	}
	if !exists {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		return
	}

	// Verify password
	if err := VerifyPassword(user.PasswordHash, req.Password); err != nil {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		return
	}

	// Generate JWT
	token, err := GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		serverError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, LoginResponse{
		Token: token,
		User:  user,
	})
}

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func handleGetUser(store *Store, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	user, found, err := store.GetUserByID(r.Context(), userID)
	if err != nil {
		serverError(w, err)
		return
	}
	if !found {
		notFound(w)
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func handleUpdateUserProfile(store *Store, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	var req UpdateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		badRequest(w, err.Error(), nil)
		return
	}

	user, err := store.UpdateUser(r.Context(), userID, req)
	if err != nil {
		serverError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func handleChangePassword(store *Store, w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	var req ChangePasswordRequest
	if err := decodeJSON(r, &req); err != nil {
		badRequest(w, err.Error(), nil)
		return
	}

	if req.CurrentPassword == "" || req.NewPassword == "" {
		badRequest(w, "Current password and new password are required", nil)
		return
	}

	if len(req.NewPassword) < 6 {
		badRequest(w, "New password must be at least 6 characters", nil)
		return
	}

	// Get user to verify current password
	user, found, err := store.GetUserByID(r.Context(), userID)
	if err != nil {
		serverError(w, err)
		return
	}
	if !found {
		notFound(w)
		return
	}

	// Verify current password
	if err := VerifyPassword(user.PasswordHash, req.CurrentPassword); err != nil {
		respondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Current password is incorrect"})
		return
	}

	// Hash new password
	newPasswordHash, err := HashPassword(req.NewPassword)
	if err != nil {
		serverError(w, err)
		return
	}

	// Update password
	if err := store.UpdateUserPassword(r.Context(), userID, newPasswordHash); err != nil {
		serverError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Password changed successfully"})
}

// corsMiddleware adds permissive CORS headers (Origin = *).
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
		w.Header().Set("Access-Control-Max-Age", "600")
		next.ServeHTTP(w, r)
	})
}

// handlePreflight returns true if the request was an OPTIONS preflight and writes the response.
func handlePreflight(w http.ResponseWriter, r *http.Request, allowed []string) bool {
	if r.Method != http.MethodOptions {
		return false
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(allowed, ","))
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")
	w.Header().Set("Access-Control-Max-Age", "600")
	w.WriteHeader(http.StatusNoContent)
	return true
}

func badRequest(w http.ResponseWriter, message string, details any) {
	respondJSON(w, http.StatusBadRequest, map[string]any{"error": message, "details": details})
}

func serverError(w http.ResponseWriter, err error) {
	log.Printf("server error: %v", err)
	respondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
}

func notFound(w http.ResponseWriter) {
	respondJSON(w, http.StatusNotFound, map[string]string{"error": "Not found"})
}

func methodNotAllowed(w http.ResponseWriter) {
	respondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
