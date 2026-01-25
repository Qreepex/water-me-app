package types

import "time"

// Upload tracks a user-owned uploaded object key in S3
type Upload struct {
	ID        string    `json:"id"        bson:"_id,omitempty"`
	UserID    string    `json:"userId"    bson:"userId"`
	Key       string    `json:"key"       bson:"key"`
	SizeBytes int64     `json:"sizeBytes" bson:"sizeBytes"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
