package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateNote(noteInput CreateNoteInput) (Note, error)
	UpdateNote(noteInput UpdateNoteInput) (Note, error)
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
		Views:            0,
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
		ID:               noteID,
		Title:            noteInput.Title,
		Tags:             noteInput.Tags,
		Content:          noteInput.Content,
		UpdatedAt:        time.Now(),
		IsStarred:        noteInput.IsStarred,
		IsCommentEnabled: noteInput.IsCommentEnabled,
		IsArchived:       true,
		Collaborators:    noteInput.Collaborators,
		Views:            0,
	}

	updatedNote, err := service.repository.UpdateNote(noteInstance)
	if err != nil {
		return updatedNote, err
	}

	return updatedNote, nil
}
