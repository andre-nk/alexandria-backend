package note

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	CreateNote(note Note) (Note, error)
	UpdateNote(note Note) (Note, error)
	DeleteNote(id string) error
	GetNoteByID(id string) (Note, error)
	GetAllNotes() ([]Note, error)
	GetNotesByUserID(uid string) ([]Note, error)
	GetFeaturedNotes(uid string) ([]Note, error)
	GetRecentNotes(uid string) ([]Note, error)
	GetStarredNotes(uid string) ([]Note, error)
	GetArchivedNotes(uid string) ([]Note, error)
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

	_, err = repo.db.Collection("notes").UpdateOne(context.Background(), bson.M{"_id": note.ID}, bson.D{{Key: "$set", Value: updateNote}})
	if err != nil {
		return note, err
	}

	return note, nil
}

func (repo *repository) DeleteNote(id string) error {
	noteID, _ := primitive.ObjectIDFromHex(id)

	_, err := repo.db.Collection("activities").DeleteMany(context.Background(), bson.M{"activity_id": noteID})
	if err != nil {
		return err
	}

	_, err = repo.db.Collection("comments").DeleteMany(context.Background(), bson.M{"note_id": noteID})
	if err != nil {
		return err
	}

	_, err = repo.db.Collection("notes").DeleteOne(
		context.Background(),
		bson.M{
			"_id": noteID,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) GetAllNotes() ([]Note, error) {
	cursor, err := repo.db.Collection("notes").Find(context.Background(), bson.M{})
	if err != nil {
		return []Note{}, nil
	}
	var notes []Note
	err = cursor.All(context.Background(), &notes)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (repo *repository) GetNoteByID(id string) (Note, error) {
	var note Note
	noteID, _ := primitive.ObjectIDFromHex(id)

	err := repo.db.Collection("notes").FindOne(
		context.Background(),
		bson.M{
			"_id": noteID,
		},
	).Decode(&note)
	if err != nil {
		return note, err
	}

	return note, nil
}

func (repo *repository) GetNotesByUserID(uid string) ([]Note, error) {
	cursor, err := repo.db.Collection("notes").Find(context.Background(), bson.M{"creator_uid": uid})
	if err != nil {
		return []Note{}, nil
	}
	var notes []Note
	err = cursor.All(context.Background(), &notes)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (repo *repository) GetFeaturedNotes(uid string) ([]Note, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "views", Value: -1}})

	cursor, err := repo.db.Collection("notes").Find(context.Background(), bson.M{"creator_uid": uid}, findOptions)
	if err != nil {
		return []Note{}, nil
	}

	var notes []Note
	err = cursor.All(context.Background(), &notes)
	if err != nil {
		return notes, err
	}

	if len(notes) >= 10 {
		notes = notes[0:10]
		return notes, nil
	}

	return notes, nil
}

func (repo *repository) GetRecentNotes(uid string) ([]Note, error) {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "updated_at", Value: -1}})

	cursor, err := repo.db.Collection("notes").Find(context.Background(), bson.M{"creator_uid": uid}, findOptions)
	if err != nil {
		return []Note{}, nil
	}

	var notes []Note
	err = cursor.All(context.Background(), &notes)
	if err != nil {
		return notes, err
	}

	if len(notes) >= 10 {
		notes = notes[0:10]
		return notes, nil
	}

	return notes, nil
}

func (repo *repository) GetStarredNotes(uid string) ([]Note, error) {
	cursor, err := repo.db.Collection("notes").Find(context.Background(), bson.M{
		"is_starred":  true,
		"creator_uid": uid,
	})
	if err != nil {
		return []Note{}, nil
	}
	var notes []Note
	err = cursor.All(context.Background(), &notes)
	if err != nil {
		return notes, err
	}

	if len(notes) >= 10 {
		notes = notes[0:10]
		return notes, nil
	}

	return notes, nil
}

func (repo *repository) GetArchivedNotes(uid string) ([]Note, error) {
	cursor, err := repo.db.Collection("notes").Find(context.Background(), bson.M{
		"is_archived": true,
		"creator_uid": uid,
	})
	if err != nil {
		return []Note{}, nil
	}
	var notes []Note
	err = cursor.All(context.Background(), &notes)
	if err != nil {
		return notes, err
	}

	if len(notes) >= 10 {
		notes = notes[0:10]
		return notes, nil
	}

	return notes, nil
}
