package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateNote(noteInput CreateNoteInput) (Note, error)
	UpdateNote(noteInput UpdateNoteInput) (Note, error)
	DeleteNote(id string) error
	GetNoteByID(id string) (Note, error)
	GetNotesByUserID(uid string) ([]Note, error)
	GetFeaturedNotes(uid string) ([]Note, error)
	GetRecentNotes(uid string) ([]Note, error)
	GetStarredNotes(uid string) ([]Note, error)
	GetArchivedNotes(uid string) ([]Note, error)
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

func (service *service) DeleteNote(id string) error {
	err := service.repository.DeleteNote(id)
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

func (service *service) GetNotesByUserID(uid string) ([]Note, error) {
	notes, err := service.repository.GetNotesByUserID(uid)
	if err != nil {
		return notes, err
	}

	return notes, err
}

func (service *service) GetFeaturedNotes(uid string) ([]Note, error) {
	notes, err := service.repository.GetFeaturedNotes(uid)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (service *service) GetRecentNotes(uid string) ([]Note, error) {
	notes, err := service.repository.GetRecentNotes(uid)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (service *service) GetStarredNotes(uid string) ([]Note, error) {
	notes, err := service.repository.GetStarredNotes(uid)
	if err != nil {
		return notes, err
	}

	return notes, nil
}

func (service *service) GetArchivedNotes(uid string) ([]Note, error) {
	notes, err := service.repository.GetArchivedNotes(uid)
	if err != nil {
		return notes, err
	}

	return notes, nil
}
