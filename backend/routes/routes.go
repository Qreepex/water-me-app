package routes

import (
	"net/http"

	"plants-backend/constants"
	"plants-backend/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, database *services.MongoDB, s3service *services.S3Service) {
	PlantHandler(router, database)
}

func getUserID(r *http.Request) (string, bool) {
	userID, ok := r.Context().Value(constants.UserIdKey).(string)
	return userID, ok
}
