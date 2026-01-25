package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionNames = []string{
	"plants",
	"uploads",
	"notifications",
}

// MongoDB wraps the MongoDB client and database
type MongoDB struct {
	client      *mongo.Client
	db          *mongo.Database
	collections map[string]*mongo.Collection
}

var mongodb *MongoDB

// Connect initializes the MongoDB client and connects to the database
func Connect(uri, db, username, password string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().
		ApplyURI(uri).
		SetServerSelectionTimeout(5 * time.Second).
		SetRetryWrites(true).SetAuth(options.Credential{
		AuthSource: db,
		Username:   username,
		Password:   password,
	})

	_client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	if err := _client.Ping(ctx, nil); err != nil {
		log.Printf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	log.Println("Connected to MongoDB")

	database := _client.Database(db)

	collections := make(map[string]*mongo.Collection)
	for _, name := range collectionNames {
		collections[name] = database.Collection(name)
	}

	mongodb = &MongoDB{
		client:      _client,
		db:          database,
		collections: collections,
	}

	return mongodb, nil
}

// GetClient returns the MongoDB client instance
func GetClient() *MongoDB {
	if mongodb == nil {
		log.Fatal("MongoDB client is not initialized. Call Connect() first.")

		return nil
	}

	return mongodb
}

// Close disconnects the MongoDB client
func (m *MongoDB) Close() {
	if err := m.client.Disconnect(context.TODO()); err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
	} else {
		log.Println("Disconnected from MongoDB")
	}
}

// GetCollection returns a MongoDB collection by its name
func (m *MongoDB) GetCollection(name string) *mongo.Collection {
	if collection, ok := m.collections[name]; ok {
		return collection
	}

	log.Fatalf("Collection %s does not exist", name)

	return nil
}
