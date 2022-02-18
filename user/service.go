package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Service interface {
	RegisterUser(user UserInput) (User, error)
	UpdateUser(user UserInput) (User, error)
	DeleteUser(id string) error
	GetUserByUID(id string) (User, error)
	GetUserByEmail(email string) (User, error)
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
		ID:          userID,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		PhotoURL:    user.PhotoURL,
		UID:         user.UID,
		Role:        user.Role,
		Location:    user.Location,
	}

	newUser, err := service.repository.RegisterUser(userInstance)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (service *service) UpdateUser(user UserInput) (User, error) {
	userInstance := User{
		UID:         user.UID,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		PhotoURL:    user.PhotoURL,
		Role:        user.Role,
		Location:    user.Location,
	}

	updatedUser, err := service.repository.UpdateUser(userInstance)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (service *service) DeleteUser(id string) error {
	err := service.repository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (service *service) GetUserByUID(id string) (User, error) {
	user, err := service.repository.GetUserByUID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (service *service) GetUserByEmail(email string) (User, error) {
	user, err := service.repository.GetUserByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}
