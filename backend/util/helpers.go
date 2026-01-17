package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// DecodeJSON decodes JSON from the request body with strict error handling.
func DecodeJSON(r *http.Request, dest any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return nil
}

// RespondJSON sends a JSON response.
func RespondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// BadRequest responds with a 400 error.
func BadRequest(w http.ResponseWriter, message string, details any) {
	RespondJSON(w, http.StatusBadRequest, map[string]any{"error": message, "details": details})
}

// ServerError responds with a 500 error.
func ServerError(w http.ResponseWriter, err error) {
	log.Printf("server error: %v", err)
	RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
}

// NotFound responds with a 404 error.
func NotFound(w http.ResponseWriter) {
	RespondJSON(w, http.StatusNotFound, map[string]string{"error": "Not found"})
}

// MethodNotAllowed responds with a 405 error.
func MethodNotAllowed(w http.ResponseWriter) {
	RespondJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
}

// Crypto helper stubs (defined in main backend for now)
func HashPassword(password string) (string, error) {
	// Implementation in main backend file
	return "", nil
}

func VerifyPassword(hash string, password string) error {
	// Implementation in main backend file
	return nil
}

func GenerateJWT(userID string, secret string) (string, error) {
	// Implementation in main backend file
	return "", nil
}
