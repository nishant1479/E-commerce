package data

import (
	"context"
	"fmt"
	"nishant/db"
	"nishant/internal/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserStore struct {
    Collection *mongo.Collection
}

// NewUserStore initializes and returns a new UserStore
func NewUserStore(db *mongo.Database) *UserStore {
    return &UserStore{
        Collection: db.Collection("users"),
    }
}


func (s *UserStore) CreateUser(ctx context.Context, user *models.User) (*models.User, error){
	res,err := s.Collection.InsertOne(ctx,user)
	if err!= nil {
		return nil,err
	}
	// get the id of the inserted documents using TYPE ASSERTION
	insertId,ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil,fmt.Errorf("failed to cast inserted ID to ObjectID")
	}
	user.Id = insertId
	return user,nil

}
