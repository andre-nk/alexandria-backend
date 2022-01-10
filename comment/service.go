package comment

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateComment(commentInput CreateCommentInput) (Comment, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) CreateComment(commentInput CreateCommentInput) (Comment, error) {
	commentInstance := Comment{
		ID:         primitive.NewObjectID(),
		NoteID:     commentInput.NoteID,
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
