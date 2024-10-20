package repository

import (
	"context"
	"shorter-url/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoURLRepository is the MongoDB implementation of the URLRepository
type MongoURLRepository struct {
	collection *mongo.Collection
}

// NewMongoURLRepository creates a new MongoDB repository
func NewMongoURLRepository(collection *mongo.Collection) *MongoURLRepository {
	return &MongoURLRepository{collection: collection}
}

// Save stores the shortened URL in MongoDB
func (m *MongoURLRepository) Save(ctx context.Context, shortURL *entity.ShortURL) error {
	_, err := m.collection.InsertOne(ctx, shortURL)
	return err
}

// FindByShortURL retrieves a URL entity based on its short URL
func (m *MongoURLRepository) FindByShortURL(ctx context.Context, shortURL string) (*entity.ShortURL, error) {
	var result entity.ShortURL
	err := m.collection.FindOne(ctx, bson.M{"shortURL": shortURL}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
