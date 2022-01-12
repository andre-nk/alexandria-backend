package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service interface {
	RegisterUser(user UserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) RegisterUser(user UserInput) (User, error) {
	userID := primitive.NewObjectID()

	userInstance := User{
		ID:       userID,
		UID:      user.UID,
		Role:     user.Role,
		Location: user.Location,
		Friends:  user.Friends,
	}

	newUser, err := service.repository.RegisterUser(userInstance)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
