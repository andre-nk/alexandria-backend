package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service interface {
	RegisterUser(user UserInput) (User, error)
	UpdateUser(user UserInput) (User, error)
	GetUserByUID(id string) (User, error)
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

func (service *service) UpdateUser(user UserInput) (User, error) {
	userInstance := User{
		UID:      user.UID,
		Role:     user.Role,
		Location: user.Location,
		Friends:  user.Friends,
	}

	updatedUser, err := service.repository.UpdateUser(userInstance)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (service *service) GetUserByUID(id string) (User, error) {
	user, err := service.repository.GetUserByUID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
