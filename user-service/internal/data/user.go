package data

import (
	"context"
	"errors"
	"fmt"
	"nishant/internal/models"

	"go.mongodb.org/mongo-driver/bson"
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

func (s*UserStore) GetUser(ctx context.Context,id string)(*models.User,error){
	var user models.User

	objectID,err :=primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil,errors.New("invalid user ID fromat")
	}
	filter :=bson.M{"_id":objectID}
	// finding user
	if err := s.Collection.FindOne(ctx,filter).Decode(&user);err !=nil {
		return nil,err
	}

	return &user,nil
}

func (s *UserStore) GetUserByEmail(ctx context.Context,email string) (*models.User,error) {
	var user models.User
	filter := bson.M{"email":email}
	if err := s.Collection.FindOne(ctx,filter).Decode(&user);err != nil{
		return nil,err
	}
	return &user,nil
}


