package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"plants-backend/middlewares"
	"plants-backend/routes"
	"plants-backend/services"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	connString := getenv("DATABASE_URL", "mongodb://localhost:27017/plants")

	db, err := services.Connect(connString, "plants", "test2", "test")
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	firebase, err := services.NewFirebaseService()
	if err != nil {
		log.Fatalf("failed to initialize firebase: %v", err)
	}

	// Protected S3 & plant routes
	s3svc, err := services.NewS3Service(ctx)
	if err != nil {
		log.Fatalf("failed to init s3: %v", err)
	}

	allowedOrigins := []string{
		"*",
	}
	cors := handlers.CORS(
		handlers.AllowedOrigins(allowedOrigins),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods(
			[]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		),
		handlers.AllowCredentials(),
	)

	r := mux.NewRouter()
	r.Use(cors)

	routes.RegisterRoutes(r, db, s3svc)

	r.Use(middlewares.AuthMiddleware(firebase))

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
