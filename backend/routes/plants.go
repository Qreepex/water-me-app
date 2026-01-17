package routes

import (
	"net/http"
	"plants-backend/services"
	"plants-backend/types"
	"plants-backend/util"
	"plants-backend/validation"

	"github.com/gorilla/mux"
)

func PlantHandler(router *mux.Router, database *services.Database) {
	router.HandleFunc("/api/plants", func(w http.ResponseWriter, r *http.Request) {
		getPlants(w, r, database)
	}).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/api/plants", func(w http.ResponseWriter, r *http.Request) {
		createPlant(w, r, database)
	}).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/api/plants/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		updatePlant(w, r, database, id)
	}).Methods(http.MethodPut, http.MethodOptions)

	router.HandleFunc("/api/plants/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		deletePlant(w, r, database, id)
	}).Methods(http.MethodDelete, http.MethodOptions)
}

func getPlants(w http.ResponseWriter, r *http.Request, db *services.Database) {
	claims := r.Context().Value("claims").(types.Claims)
	userID := claims.ID

	plants, err := db.ListPlants(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to retrieve plants", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, 200, plants)
}

func createPlant(w http.ResponseWriter, r *http.Request, db *services.Database) {
	claims := r.Context().Value("claims").(types.Claims)
	userID := claims.ID

	var input types.PlantInput
	if err := util.DecodeJSON(r, &input); err != nil {
		util.BadRequest(w, err.Error(), nil)
		return
	}

	sanitized := validation.SanitizePlantInput(input)
	errors := validation.ValidatePlantInput(sanitized, false)
	if len(errors) > 0 {
		util.BadRequest(w, "Validation failed", errors)
		return
	}

	plant, err := db.CreatePlant(r.Context(), userID, sanitized)
	if err != nil {
		util.ServerError(w, err)
		return
	}
	util.RespondJSON(w, http.StatusCreated, plant)
}

func updatePlant(w http.ResponseWriter, r *http.Request, db *services.Database, id string) {
	userID, ok := getUserID(r)
	if !ok {
		util.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	var input types.PlantInput
	if err := util.DecodeJSON(r, &input); err != nil {
		util.BadRequest(w, err.Error(), nil)
		return
	}

	sanitized := validation.SanitizePlantInput(input)
	errors := validation.ValidatePlantInput(sanitized, false)
	if len(errors) > 0 {
		util.BadRequest(w, "Validation failed", errors)
		return
	}

	plant, found, err := db.UpdatePlant(r.Context(), id, userID, sanitized)
	if err != nil {
		util.ServerError(w, err)
		return
	}
	if !found {
		util.NotFound(w)
		return
	}
	util.RespondJSON(w, http.StatusOK, plant)
}

func deletePlant(w http.ResponseWriter, r *http.Request, db *services.Database, id string) {
	userID, ok := getUserID(r)
	if !ok {
		util.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		return
	}

	deleted, err := db.DeletePlant(r.Context(), id, userID)
	if err != nil {
		util.ServerError(w, err)
		return
	}

	if !deleted {
		util.NotFound(w)
		return
	}
	util.RespondJSON(w, http.StatusOK, map[string]bool{"success": true})
}
