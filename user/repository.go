package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	RegisterUser(user User) (User, error)
	GetUserByUID(id string) (User, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
}

func (repo *repository) RegisterUser(user User) (User, error) {
	_, err := repo.db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *repository) GetUserByUID(id string) (User, error) {
	var user User

	err := repo.db.Collection("users").FindOne(
		context.Background(),
		bson.M{
			"uid": id,
		},
	).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
