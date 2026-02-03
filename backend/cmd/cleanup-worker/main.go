package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/qreepex/water-me-app/backend/services"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()
	connString := getenv("DATABASE_URL", "mongodb://localhost:27017/plants")
	mongoUser := getenv("MONGODB_USERNAME", "test2")
	mongoPassword := getenv("MONGODB_PASSWORD", "test")
	mongoDatabase := getenv("MONGODB_DATABASE", "plants")

	db, err := services.Connect(connString, mongoDatabase, mongoUser, mongoPassword)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	s3svc, err := services.NewS3Service(ctx)
	if err != nil {
		log.Fatalf("failed to init s3: %v", err)
	}

	runCleanupCheck(db, s3svc)
}

// runCleanupCheck runs a background job to clean up orphaned uploads every 30 minutes
func runCleanupCheck(db *services.MongoDB, s3 *services.S3Service) {
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	log.Println("Orphaned upload cleanup worker started (runs every 30 minutes)")

	for range ticker.C {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		uploadSvc := services.NewUploadService(db, s3)
		count, err := uploadSvc.CleanupOrphanedUploads(ctx, 1*time.Hour)
		cancel()

		if err != nil {
			log.Printf("Cleanup worker error: %v", err)
		} else if count > 0 {
			log.Printf("Cleaned up %d orphaned uploads", count)
		}
	}
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
