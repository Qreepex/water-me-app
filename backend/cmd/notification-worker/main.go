package main

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/qreepex/water-me-app/backend/services"

	_ "github.com/joho/godotenv/autoload"
)

const (
	plantsBatchSize = 1000 // Process 1000 plants per query
	workerInterval  = 5 * time.Minute
)

func main() {
	ctx := context.Background()

	// Load configuration from environment
	config := loadConfig()

	// Initialize database connection
	db, err := services.Connect(
		config.DatabaseURL,
		config.DatabaseName,
		config.DatabaseUser,
		config.DatabasePassword,
	)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize Firebase service
	firebase, err := services.NewFirebaseService()
	if err != nil {
		log.Fatalf("failed to initialize firebase: %v", err)
	}

	// Load notification messages from JSON files
	messageDir := filepath.Join(".", "messages")
	if err := services.LoadNotificationMessages(messageDir); err != nil {
		log.Fatalf("failed to load notification messages: %v", err)
	}

	log.Println("========================================")
	log.Println("Notification Worker Started")
	log.Println("========================================")
	log.Printf("Scalable for 50k+ users with 100+ plants each")
	log.Printf("Configuration:")
	log.Printf("  - Plants batch size: %d", plantsBatchSize)
	log.Printf("  - FCM batch size: %d", services.FCMBatchSize)
	log.Printf("  - Check interval: %v", workerInterval)
	log.Printf("  - Cooldown period: %v", services.NotificationCooldown)
	log.Println("========================================")

	// Run notification check loop
	ticker := time.NewTicker(workerInterval)
	defer ticker.Stop()

	// Run immediately on startup
	runNotificationCheck(ctx, db, firebase)

	// Then run every interval
	for range ticker.C {
		runNotificationCheck(ctx, db, firebase)
	}
}

// runNotificationCheck orchestrates the entire notification check process
func runNotificationCheck(
	ctx context.Context,
	db *services.MongoDB,
	firebase *services.FirebaseService,
) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Minute) // Long timeout for large batches
	defer cancel()

	startTime := time.Now()
	log.Println("========== Starting notification check ==========")

	// Process all notification types and get statistics
	stats := services.ProcessNotifications(ctx, db, firebase, plantsBatchSize)

	// Log results
	duration := time.Since(startTime)
	log.Println("========== Notification check complete ==========")
	log.Printf("Duration: %v", duration)
	log.Printf("Plants checked: %d", stats.PlantsChecked)
	log.Printf("Notifications sent: %d", stats.NotificationsSent)
	log.Printf("Notifications failed: %d", stats.NotificationsFailed)
	log.Printf("Users notified: %d", stats.UsersNotified)
	if stats.NotificationsSent > 0 {
		successRate := float64(
			stats.NotificationsSent,
		) / float64(
			stats.NotificationsSent+stats.NotificationsFailed,
		) * 100
		log.Printf("Success rate: %.2f%%", successRate)
	}
	log.Println("================================================")
}

type Config struct {
	DatabaseURL      string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func loadConfig() Config {
	return Config{
		DatabaseURL:      getenv("DATABASE_URL", "mongodb://localhost:27017/plants"),
		DatabaseUser:     getenv("MONGODB_USERNAME", "test2"),
		DatabasePassword: getenv("MONGODB_PASSWORD", "test"),
		DatabaseName:     getenv("MONGODB_DATABASE", "plants"),
	}
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
