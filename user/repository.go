package user

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	RegisterUser(user User) (User, error)
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
