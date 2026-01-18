package routes

import (
	"log"
	"net/http"
	"plants-backend/services"
	"plants-backend/types"
	"plants-backend/util"
	"plants-backend/validation"
	"time"

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

	// router.HandleFunc("/api/plants/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	id := vars["id"]
	// 	updatePlant(w, r, database, id)
	// }).Methods(http.MethodPut, http.MethodOptions)

	// router.HandleFunc("/api/plants/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	id := vars["id"]
	// 	deletePlant(w, r, database, id)
	// }).Methods(http.MethodDelete, http.MethodOptions)
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

// func updatePlant(w http.ResponseWriter, r *http.Request, db *services.MongoDB, id string) {
// 	userID, ok := getUserID(r)
// 	if !ok {
// 		util.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
// 		return
// 	}

// 	var input types.PlantInput
// 	if err := util.DecodeJSON(r, &input); err != nil {
// 		util.BadRequest(w, err.Error(), nil)
// 		return
// 	}

// 	sanitized := validation.SanitizePlantInput(input)
// 	errors := validation.ValidatePlantInput(sanitized, false)
// 	if len(errors) > 0 {
// 		util.BadRequest(w, "Validation failed", errors)
// 		return
// 	}

// 	plant, found, err := db.UpdatePlant(r.Context(), id, userID, sanitized)
// 	if err != nil {
// 		util.ServerError(w, err)
// 		return
// 	}
// 	if !found {
// 		util.NotFound(w)
// 		return
// 	}
// 	util.RespondJSON(w, http.StatusOK, plant)
// }

// func deletePlant(w http.ResponseWriter, r *http.Request, db *services.Database, id string) {
// 	userID, ok := getUserID(r)
// 	if !ok {
// 		util.RespondJSON(w, http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
// 		return
// 	}

// 	deleted, err := db.DeletePlant(r.Context(), id, userID)
// 	if err != nil {
// 		util.ServerError(w, err)
// 		return
// 	}

// 	if !deleted {
// 		util.NotFound(w)
// 		return
// 	}
// 	util.RespondJSON(w, http.StatusOK, map[string]bool{"success": true})
// }

func plantInputToPlant(input types.PlantInput, userID string) types.Plant {
	return types.Plant{
		UserID:                  userID,
		Name:                    *input.Name,
		Species:                 input.Species,
		SunLight:                input.SunLight,
		WateringIntervalDays:    input.WateringIntervalDays,
		LastWatered:             parseTimePointer(input.LastWatered),
		Notes:                   *input.Notes,
		Flags:                   *input.Flags,
		FertilizingIntervalDays: input.FertilizingIntervalDays,
		LastFertilized:          parseTimePointer(input.LastFertilized),
		PhotoIDs:                *input.PhotoIDs,
	}
}

func createPlantFromRequest(req types.CreatePlantRequest, userID string) types.Plant {
	now := time.Now()
	id, err := gonanoid.New()
	if err != nil {
		log.Printf("Failed to generate plant ID: %v", err)
		id = ""
	}

	plant := types.Plant{
		ID:                      id,
		UserID:                  userID,
		Name:                    req.Name,
		Species:                 req.Species,
		SunLight:                req.SunLight,
		WateringIntervalDays:    req.WateringIntervalDays,
		LastWatered:             parseTimePointer(req.LastWatered),
		FertilizingIntervalDays: req.FertilizingIntervalDays,
		LastFertilized:          parseTimePointer(req.LastFertilized),
		PreferedTemperature:     req.PreferedTemperature,
		PreferedHumidity:        req.PreferedHumidity,
		SprayIntervalDays:       req.SprayIntervalDays,
		CreatedAt:               now,
		UpdatedAt:               now,
	}

	if req.Notes != nil {
		plant.Notes = *req.Notes
	}
	if req.Flags != nil {
		plant.Flags = *req.Flags
	}
	if req.PhotoIDs != nil {
		plant.PhotoIDs = *req.PhotoIDs
	}

	return plant
}

func parseTimePointer(timeStr *string) *time.Time {
	if timeStr == nil {
		return nil
	}
	t, err := time.Parse(time.RFC3339Nano, *timeStr)
	if err != nil {
		return nil
	}
	return &t
}
