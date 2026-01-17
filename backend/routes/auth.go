package routes

import (
	"net/http"
	"plants-backend/services"
	"plants-backend/types"
	"plants-backend/util"

	"github.com/gorilla/mux"
)

func AuthHandler(router *mux.Router, database *services.Database) {
	router.HandleFunc("/api/signup", func(w http.ResponseWriter, r *http.Request) {
		handleSignup(w, r, database)
	}).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		handleLogin(w, r, database)
	}).Methods(http.MethodPost, http.MethodOptions)
}

func handleSignup(w http.ResponseWriter, r *http.Request, db *services.Database) {
	var req types.SignupRequest
	if err := util.DecodeJSON(r, &req); err != nil {
		util.BadRequest(w, err.Error(), nil)
		return
	}

	if req.Email == "" || req.Password == "" {
		util.BadRequest(w, "Email and password are required", nil)
		return
	}

	if len(req.Password) < 6 {
		util.BadRequest(w, "Password must be at least 6 characters", nil)
		return
	}

	// Check if user already exists
	_, exists, err := db.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		util.ServerError(w, err)
		return
	}
	if exists {
		util.BadRequest(w, "User with this email already exists", nil)
		return
	}

	// Hash password
	passwordHash, err := util.HashPassword(req.Password)
	if err != nil {
		util.ServerError(w, err)
		return
	}

	// Create user
	user, err := db.CreateUser(r.Context(), req.Email, passwordHash)
	if err != nil {
		util.ServerError(w, err)
		return
	}

	// Generate JWT
	token, err := util.GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		util.ServerError(w, err)
		return
	}

	util.RespondJSON(w, http.StatusCreated, types.LoginResponse{
		Token: token,
		User:  user,
	})
}

func handleLogin(w http.ResponseWriter, r *http.Request, db *services.Database) {
	var req types.LoginRequest
	if err := util.DecodeJSON(r, &req); err != nil {
		util.BadRequest(w, err.Error(), nil)
		return
	}

	if req.Email == "" || req.Password == "" {
		util.BadRequest(w, "Email and password are required", nil)
		return
	}

	// Get user by email
	user, exists, err := db.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		util.ServerError(w, err)
		return
	}
	if !exists {
		util.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		return
	}

	// Verify password
	if err := util.VerifyPassword(user.PasswordHash, req.Password); err != nil {
		util.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
		return
	}

	// Generate JWT
	token, err := util.GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		util.ServerError(w, err)
		return
	}

	util.RespondJSON(w, http.StatusOK, types.LoginResponse{
		Token: token,
		User:  user,
	})
}
