package routes

import (
	"net/http"
	"plants-backend/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, database *services.Database, s3service services.S3Service) {
	PlantHandler(router, database)
	UserHandler(router, database)
	AuthHandler(router, database)
}

func getUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(userIDKey).(string)
	return userID, ok
}
