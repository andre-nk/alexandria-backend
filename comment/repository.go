package comment

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateComment(comment Comment) (Comment, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
}

func (repo *repository) CreateComment(comment Comment) (Comment, error) {
	_, err := repo.db.Collection("comments").InsertOne(context.Background(), comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
