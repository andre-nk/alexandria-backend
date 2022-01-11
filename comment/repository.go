package comment

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateComment(comment Comment) (Comment, error)
	GetCommentsByNoteID(id string) ([]Comment, error)
	GetCommentByID(id string) (Comment, error)
	DeleteComment(id string) error
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

func (repo *repository) GetCommentsByNoteID(id string) ([]Comment, error) {
	noteID, _ := primitive.ObjectIDFromHex(id)

	cursor, err := repo.db.Collection("comments").Find(context.Background(), bson.M{"note_id": noteID})
	if err != nil {
		return []Comment{}, nil
	}
	var comments []Comment
	err = cursor.All(context.Background(), &comments)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (repo *repository) GetCommentByID(id string) (Comment, error) {
	var comment Comment
	commentID, _ := primitive.ObjectIDFromHex(id)

	err := repo.db.Collection("comments").FindOne(
		context.Background(),
		bson.M{
			"_id": commentID,
		},
	).Decode(&comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (repo *repository) DeleteComment(id string) error {
	commentID, _ := primitive.ObjectIDFromHex(id)

	_, err := repo.db.Collection("comments").DeleteOne(
		context.Background(),
		bson.M{
			"_id": commentID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
