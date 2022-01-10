package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateNote(noteInput CreateNoteInput) (Note, error)
	UpdateNote(noteInput UpdateNoteInput) (Note, error)
	DeleteNote(noteInput UpdateNoteInput) error
	GetNoteByID(id string) (Note, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) CreateNote(noteInput CreateNoteInput) (Note, error) {
	noteInstance := Note{
		ID:               primitive.NewObjectID(),
		Title:            noteInput.Title,
		CreatorUID:       noteInput.CreatorUID,
		Tags:             noteInput.Tags,
		Content:          noteInput.Content,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		IsStarred:        noteInput.IsStarred,
		IsCommentEnabled: noteInput.IsCommentEnabled,
		IsArchived:       false,
		Collaborators:    noteInput.Collaborators,
		Views:            1,
	}

	newNote, err := service.repository.CreateNote(noteInstance)
	if err != nil {
		return newNote, err
	}

	return newNote, nil
}

func (service *service) UpdateNote(noteInput UpdateNoteInput) (Note, error) {
	noteID, _ := primitive.ObjectIDFromHex(noteInput.ID.ID)

	noteInstance := Note{
		ID:               noteID, //for the sake of response only, not to be inserted
		Title:            noteInput.Title,
		Tags:             noteInput.Tags,
		Content:          noteInput.Content,
		UpdatedAt:        time.Now(),
		IsStarred:        noteInput.IsStarred, //bool not updated?
		IsCommentEnabled: noteInput.IsCommentEnabled,
		IsArchived:       noteInput.IsArchived,
		Collaborators:    noteInput.Collaborators,
		Views:            noteInput.Views,
	}

	updatedNote, err := service.repository.UpdateNote(noteInstance)
	if err != nil {
		return updatedNote, err
	}

	return updatedNote, nil
}

func (service *service) DeleteNote(noteInput UpdateNoteInput) error {
	noteID, _ := primitive.ObjectIDFromHex(noteInput.ID.ID)

	noteInstance := Note{
		ID: noteID,
	}

	err := service.repository.DeleteNote(noteInstance)
	if err != nil {
		return err
	}

	return nil
}

func (service *service) GetNoteByID(id string) (Note, error) {
	note, err := service.repository.GetNoteByID(id)
	if err != nil {
		return note, err
	}

	return note, nil
}
