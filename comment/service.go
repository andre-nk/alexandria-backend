package comment

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateComment(commentInput CreateCommentInput) (Comment, error)
	GetCommentsByNoteID(id string) ([]Comment, error)
	GetCommentByID(id string) (Comment, error)
	DeleteComment(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) CreateComment(commentInput CreateCommentInput) (Comment, error) {
	noteID, err := primitive.ObjectIDFromHex(commentInput.NoteID)
	if err != nil {
		return Comment{}, nil
	}

	commentInstance := Comment{
		ID:         primitive.NewObjectID(),
		NoteID:     noteID,
		CreatorUID: commentInput.CreatorUID,
		CreatedAt:  time.Now(),
		Content:    commentInput.Content,
		Mentions:   commentInput.Mentions,
	}

	newComment, err := service.repository.CreateComment(commentInstance)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

func (service *service) GetCommentsByNoteID(id string) ([]Comment, error) {
	comments, err := service.repository.GetCommentsByNoteID(id)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (service *service) GetCommentByID(id string) (Comment, error) {
	comment, err := service.repository.GetCommentByID(id)
	if err != nil {
		return comment, err
	}

	return comment, nil
}

func (service *service) DeleteComment(id string) error {
	err := service.repository.DeleteComment(id)
	if err != nil {
		return err
	}

	return nil
}
