package services

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Upload represents a temporary uploaded file waiting to be attached to a plant
type Upload struct {
	ID        string    `bson:"_id"`
	Key       string    `bson:"key"`
	UserID    string    `bson:"userId"`
	CreatedAt time.Time `bson:"createdAt"`
}

// UploadService manages upload records and cleanup
type UploadService struct {
	collection *mongo.Collection
	s3Service  *S3Service
}

// NewUploadService creates a new upload service
func NewUploadService(db *MongoDB, s3Service *S3Service) *UploadService {
	return &UploadService{
		collection: db.GetCollection("uploads"),
		s3Service:  s3Service,
	}
}

// RegisterUpload records an uploaded file in the database
func (us *UploadService) RegisterUpload(ctx context.Context, key, userID string) (string, error) {
	upload := Upload{
		ID:        key, // Use key as ID for easy reference
		Key:       key,
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	_, err := us.collection.InsertOne(ctx, upload)
	if err != nil {
		return "", fmt.Errorf("register upload: %w", err)
	}

	return key, nil
}

// DeleteUpload removes an upload record and deletes the file from S3
func (us *UploadService) DeleteUpload(ctx context.Context, key, userID string) error {
	// Verify the key belongs to this user before deleting
	if !KeyBelongsToUser(key, userID) {
		return fmt.Errorf("unauthorized: key does not belong to user")
	}

	// Delete from S3
	if err := us.s3Service.DeleteObject(ctx, key); err != nil {
		return fmt.Errorf("delete from s3: %w", err)
	}

	// Delete from database
	_, err := us.collection.DeleteOne(ctx, bson.M{"_id": key})
	if err != nil {
		return fmt.Errorf("delete upload record: %w", err)
	}

	return nil
}

// CleanupOrphanedUploads deletes uploads older than maxAge that haven't been applied to a plant
func (us *UploadService) CleanupOrphanedUploads(
	ctx context.Context,
	maxAge time.Duration,
) (int, error) {
	cutoffTime := time.Now().Add(-maxAge)

	// Find all uploads older than maxAge
	cursor, err := us.collection.Find(ctx, bson.M{
		"createdAt": bson.M{"$lt": cutoffTime},
	})
	if err != nil {
		return 0, fmt.Errorf("query orphaned uploads: %w", err)
	}
	defer cursor.Close(ctx)

	var uploads []Upload
	if err = cursor.All(ctx, &uploads); err != nil {
		return 0, fmt.Errorf("decode uploads: %w", err)
	}

	if len(uploads) == 0 {
		return 0, nil
	}

	// Filter to only unreferenced uploads
	var keysToDelete []string
	var idsToDelete []interface{}

	for _, u := range uploads {
		// Check if this upload is referenced in any plant
		referenced, err := us.IsUploadReferenced(ctx, u.Key)
		if err != nil {
			// Log error but continue checking others
			fmt.Printf("Error checking reference for %s: %v\n", u.Key, err)
			continue
		}

		// Only delete if NOT referenced
		if !referenced {
			keysToDelete = append(keysToDelete, u.Key)
			idsToDelete = append(idsToDelete, u.ID)
		}
	}

	if len(keysToDelete) == 0 {
		return 0, nil
	}

	// Delete from S3
	if err := us.s3Service.DeleteObjects(ctx, keysToDelete); err != nil {
		return 0, fmt.Errorf("batch delete from s3: %w", err)
	}

	// Delete from database
	result, err := us.collection.DeleteMany(ctx, bson.M{
		"_id": bson.M{"$in": idsToDelete},
	})
	if err != nil {
		return 0, fmt.Errorf("batch delete uploads: %w", err)
	}

	return int(result.DeletedCount), nil
}

// IsUploadReferenced checks if an upload is referenced in any plant
func (us *UploadService) IsUploadReferenced(ctx context.Context, key string) (bool, error) {
	plantsCollection := us.collection.Database().Collection("plants")
	count, err := plantsCollection.CountDocuments(ctx, bson.M{
		"photoIds": key,
	})
	if err != nil {
		return false, fmt.Errorf("check reference: %w", err)
	}
	return count > 0, nil
}
