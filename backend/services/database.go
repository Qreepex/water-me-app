package services

import (
	"context"
	"plants-backend/constants"
	"plants-backend/types"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) GetPlants(ctx context.Context, userID string) ([]types.Plant, error) {
	collection := m.GetCollection(constants.MongoDBCollections.Plants)
	if collection == nil {
		return nil, types.ErrNoDocuments
	}

	var plants []types.Plant
	cursor, err := collection.Find(ctx, bson.M{"userId": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &plants)
	if err != nil {
		return nil, err
	}

	return plants, nil
}

func (m *MongoDB) CreatePlant(ctx context.Context, plantInput types.Plant) (*types.Plant, error) {
	collection := m.GetCollection(constants.MongoDBCollections.Plants)
	if collection == nil {
		return nil, types.ErrNoDocuments
	}

	_, err := collection.InsertOne(ctx, plantInput)
	return &plantInput, err
}
