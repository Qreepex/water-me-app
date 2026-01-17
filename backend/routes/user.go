package routes

import (
	"net/http"
	"plants-backend/services"
	"plants-backend/types"
	"plants-backend/util"

	"github.com/gorilla/mux"
)

func UserHandler(router *mux.Router, database *services.Database) {
	router.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		getUser(w, r, database)
	}).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		updateUserProfile(w, r, database)
	}).Methods(http.MethodPut, http.MethodOptions)
}

func getUser(w http.ResponseWriter, r *http.Request, db *services.Database) {
	claims := r.Context().Value("claims").(types.Claims)
	userID := claims.ID

	user, found, err := db.GetUserByID(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	if !found {
		util.NotFound(w)
		return
	}

	util.RespondJSON(w, http.StatusOK, user)
}

func updateUserProfile(w http.ResponseWriter, r *http.Request, db *services.Database) {
	claims := r.Context().Value("claims").(types.Claims)
	userID := claims.ID

	var input types.UpdateUserRequest
	if err := util.DecodeJSON(r, &input); err != nil {
		util.BadRequest(w, err.Error(), nil)
		return
	}

	// TODO: add user profile update validation
	sanitized := validation.SanitizeUserProfileUpdateInput(input)
	errors := validation.ValidateUserProfileUpdateInput(sanitized)

	if len(errors) > 0 {
		util.BadRequest(w, "Validation errors", errors)
		return
	}

	updatedUser, err := db.UpdateUser(r.Context(), userID, sanitized)
	if err != nil {
		util.ServerError(w, err)
		return
	}

	util.RespondJSON(w, http.StatusOK, updatedUser)
}
