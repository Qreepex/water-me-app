package routes

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
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

	router.HandleFunc("/api/plants/slug/{slug}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		slug := vars["slug"]
		getPlantBySlug(w, r, database, slug)
	}).Methods(http.MethodGet, http.MethodOptions)

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

func getPlantBySlug(w http.ResponseWriter, r *http.Request, db *services.MongoDB, slug string) {
	userID, ok := getUserID(r)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.Printf("Getting plant by slug '%s' for user %v", slug, userID)

	plant, err := db.GetPlantBySlug(r.Context(), userID, slug)
	if err != nil {
		log.Printf("Failed to retrieve plant by slug: %v", err)
		http.Error(w, "Failed to retrieve plant", http.StatusInternalServerError)
		return
	}

	if plant == nil {
		util.NotFound(w)
		return
	}

	util.RespondJSON(w, http.StatusOK, plant)
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

	// Get existing plants to check for slug uniqueness
	existingPlants, err := db.GetPlants(r.Context(), userID)
	if err != nil {
		util.ServerError(w, err)
		return
	}

	plant := createPlantFromRequest(req, userID, existingPlants)
	createdPlant, err := db.CreatePlant(r.Context(), plant)
	if err != nil {
		util.ServerError(w, err)
		return
	}
	util.RespondJSON(w, http.StatusCreated, createdPlant)
}

// slugify converts a string to a URL-friendly slug
func slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)
	// Replace spaces and special characters with hyphens
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")
	// Remove consecutive hyphens
	reg = regexp.MustCompile(`-+`)
	s = reg.ReplaceAllString(s, "-")
	// Trim hyphens from start and end
	s = strings.Trim(s, "-")
	return s
}

// generateUniqueSlug creates a unique slug for the plant within the user's collection
func generateUniqueSlug(name string, location types.Location, existingPlants []types.Plant) string {
	baseSlug := slugify(name)
	if baseSlug == "" {
		baseSlug = "plant"
	}

	// Check if base slug is unique
	if !slugExists(baseSlug, existingPlants) {
		return baseSlug
	}

	// Try with location
	locationPart := slugify(location.Room)
	if locationPart == "" {
		locationPart = slugify(location.Position)
	}
	if locationPart != "" {
		slugWithLocation := baseSlug + "-" + locationPart
		if !slugExists(slugWithLocation, existingPlants) {
			return slugWithLocation
		}
	}

	// Add number suffix
	counter := 1
	for {
		numberedSlug := fmt.Sprintf("%s-%d", baseSlug, counter)
		if !slugExists(numberedSlug, existingPlants) {
			return numberedSlug
		}
		counter++
	}
}

// slugExists checks if a slug already exists in the user's plants
func slugExists(slug string, plants []types.Plant) bool {
	for _, plant := range plants {
		if plant.Slug == slug {
			return true
		}
	}
	return false
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

func createPlantFromRequest(
	req types.CreatePlantRequest,
	userID string,
	existingPlants []types.Plant,
) types.Plant {
	now := time.Now()
	id, err := gonanoid.New()
	if err != nil {
		log.Printf("Failed to generate plant ID: %v", err)
		id = ""
	}

	// Generate unique slug
	slug := generateUniqueSlug(req.Name, req.Location, existingPlants)

	plant := types.Plant{
		ID:                  id,
		UserID:              userID,
		Slug:                slug,
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
