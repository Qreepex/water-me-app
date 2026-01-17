package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"plants-backend/middlewares"
	"plants-backend/services"
	"plants-backend/types"
	"plants-backend/validation"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	connString := getenv("DATABASE_URL", "postgres://postgres:pw@localhost:5432/plants?sslmode=disable")
	jwtSecret := getenv("JWT_SECRET", "your-secret-key-change-this-in-production")

	db, err := services.NewDatabase(ctx, connString)
	if err != nil {
		log.Fatalf("failed to initialize store: %v", err)
	}
	defer db.Close()
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

	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/api/signup", func(w http.ResponseWriter, r *http.Request) {
		routes.HandleSignup(store, jwtSecret, w, r)
	})
	mux.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		routes.HandleLogin(store, jwtSecret, w, r)
	})

	// Protected user routes
	mux.HandleFunc("/api/user", middlewares.AuthMiddleware(jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			routes.HandleGetUser(store, w, r)
		case http.MethodPut:
			routes.HandleUpdateUserProfile(store, w, r)
		default:
			routes.MethodNotAllowed(w)
		}
	}))

	mux.HandleFunc("/api/user/password", middlewares.AuthMiddleware(jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		routes.HandleChangePassword(store, w, r)
	}))

	// Protected S3 & plant routes
	s3svc, err := services.NewS3Service(ctx)
	if err != nil {
		log.Fatalf("failed to init s3: %v", err)
	}

	// Uploads router (presign)
	mux.Handle("/api/uploads/", middlewares.AuthMiddleware(jwtSecret, func(w http.ResponseWriter, r *http.Request) {
		routes.UploadsRouter(s3svc).ServeHTTP(w, r)
	}))

	// Plants router (CRUD with S3 integration)
	plantsHandler := routes.CreatePlantsHandler(db, s3svc,
		func(input any) any { return validation.SanitizePlantInput(input.(types.PlantInput)) },
		func(input any, partial bool) []types.ValidationError {
			return validation.ValidatePlantInput(input.(types.PlantInput), partial)
		},
	)
	mux.Handle("/api/plants", middlewares.AuthMiddleware(jwtSecret, plantsHandler))
	mux.Handle("/api/plants/", middlewares.AuthMiddleware(jwtSecret, plantsHandler))

	addr := getenv("PORT", "8080")
	srv := &http.Server{
		Addr:         ":" + addr,
		Handler:      middlewares.CORS(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
