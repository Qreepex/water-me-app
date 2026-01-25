package routes

import (
	"log"
	"net/http"
	"time"

	"plants-backend/services"
	"plants-backend/types"
	"plants-backend/util"
	"plants-backend/validation"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/gorilla/mux"
)

func PlantHandler(router *mux.Router, database *services.MongoDB) {
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
	}).Methods(http.MethodPatch, http.MethodOptions)

	router.HandleFunc("/api/plants/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		deletePlant(w, r, database, id)
	}).Methods(http.MethodDelete, http.MethodOptions)
}

func getPlants(w http.ResponseWriter, r *http.Request, db *services.MongoDB) {
	userID, ok := getUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("Getting plants for %v", userID)

	plants, err := db.GetPlants(r.Context(), userID)
	if err != nil {
		log.Printf("Failed to retrieve plants: %v", err)

		http.Error(w, "Failed to retrieve plants", http.StatusInternalServerError)
		return
	}

	// Return empty array instead of null if no plants exist
	if plants == nil {
		plants = []types.Plant{}
	}

	util.RespondJSON(w, 200, plants)
}

func createPlant(w http.ResponseWriter, r *http.Request, db *services.MongoDB) {
	userID, ok := getUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req types.CreatePlantRequest
	if err := util.DecodeJSON(r, &req); err != nil {
		util.BadRequest(w, err.Error(), nil)
		return
	}

	errors := validation.ValidateCreatePlantRequest(req)
	if len(errors) > 0 {
		util.BadRequest(w, "Validation failed", errors)
		return
	}

	plant := createPlantFromRequest(req, userID)
	createdPlant, err := db.CreatePlant(r.Context(), plant)
	if err != nil {
		util.ServerError(w, err)
		return
	}
	util.RespondJSON(w, http.StatusCreated, createdPlant)
}

func updatePlant(w http.ResponseWriter, r *http.Request, db *services.MongoDB, id string) {
	userID, ok := getUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req types.UpdatePlantRequest
	if err := util.DecodeJSON(r, &req); err != nil {
		util.BadRequest(w, err.Error(), nil)
		return
	}

	errors := validation.ValidateUpdatePlantRequest(req)
	if len(errors) > 0 {
		util.BadRequest(w, "Validation failed", errors)
		return
	}

	plant, found, err := db.UpdatePlant(r.Context(), id, userID, req)
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

func deletePlant(w http.ResponseWriter, r *http.Request, db *services.MongoDB, id string) {
	userID, ok := getUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
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

func createPlantFromRequest(req types.CreatePlantRequest, userID string) types.Plant {
	now := time.Now()
	id, err := gonanoid.New()
	if err != nil {
		log.Printf("Failed to generate plant ID: %v", err)
		id = ""
	}

	plant := types.Plant{
		ID:                  id,
		UserID:              userID,
		Name:                req.Name,
		Species:             req.Species,
		IsToxic:             req.IsToxic,
		Sunlight:            req.Sunlight,
		PreferedTemperature: req.PreferedTemperature,
		Location:            req.Location,
		Watering:            req.Watering,
		Fertilizing:         req.Fertilizing,
		Humidity:            req.Humidity,
		Soil:                req.Soil,
		Seasonality:         req.Seasonality,
		PestHistory:         req.PestHistory,
		Flags:               req.Flags,
		Notes:               req.Notes,
		PhotoIDs:            req.PhotoIDs,
		GrowthHistory:       req.GrowthHistory,
		CreatedAt:           now,
		UpdatedAt:           now,
	}

	return plant
}
