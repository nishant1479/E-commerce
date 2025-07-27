package data

import (
	"context"
	"errors"
	"nishant/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductStore struct {
	Collection *mongo.Collection
}

func NewProductStore(db *mongo.Database) *ProductStore {
	return &ProductStore{
		Collection: db.Collection("products"),
	}
}
func (p *ProductStore) GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product
	cursor, err := p.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var product *model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductStore) GetProductById(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product

	// Convert string ID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	filter := bson.M{"_id": objectID}
	if err := p.Collection.FindOne(ctx, filter).Decode(&product); err != nil {
		return &product, err
	}
	return &product, nil
}