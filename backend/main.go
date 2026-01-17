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

	store, err := NewStore(ctx, connString)
	if err != nil {
		log.Fatalf("failed to initialize store: %v", err)
	}
	defer store.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/plants", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleListPlants(store, w, r)
		case http.MethodPost:
			handleCreatePlant(store, w, r)
		default:
			methodNotAllowed(w)
		}
	})

	mux.HandleFunc("/api/plants/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/api/plants/")
		if id == "" || strings.Contains(id, "/") {
			notFound(w)
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
	})

	addr := getenv("PORT", "8080")
	srv := &http.Server{
		Addr:         ":" + addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func handleListPlants(store *Store, w http.ResponseWriter, r *http.Request) {
	plants, err := store.ListPlants(r.Context())
	if err != nil {
		serverError(w, err)
		return
	}
	respondJSON(w, http.StatusOK, plants)
}

func handleGetPlant(store *Store, id string, w http.ResponseWriter, r *http.Request) {
	plant, found, err := store.GetPlant(r.Context(), id)
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

	plant, err := store.CreatePlant(r.Context(), sanitized)
	if err != nil {
		serverError(w, err)
		return
	}
	respondJSON(w, http.StatusCreated, plant)
}

func handleUpdatePlant(store *Store, id string, partial bool, w http.ResponseWriter, r *http.Request) {
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

	plant, found, err := store.UpdatePlant(r.Context(), id, sanitized)
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
	deleted, err := store.DeletePlant(r.Context(), id)
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

func respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
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
