package note

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateNote(note Note) (Note, error)
	UpdateNote(note Note) (Note, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
}

func (repo *repository) CreateNote(note Note) (Note, error) {
	_, err := repo.db.Collection("notes").InsertOne(context.Background(), note)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (repo *repository) UpdateNote(note Note) (Note, error) {
	noteByte, err := bson.Marshal(note)
	if err != nil {
		return note, err
	}

	var updateNote bson.M
	err = bson.Unmarshal(noteByte, &updateNote)
	if err != nil {
		return note, err
	}

	_, err = repo.db.Collection("notes").UpdateOne(context.Background(), bson.M{"_id": note.ID}, bson.D{{"$set", updateNote}})
	if err != nil {
		return note, err
	}

	return note, nil
}
