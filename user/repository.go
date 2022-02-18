package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	RegisterUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id string) error
	GetUserByUID(id string) (User, error)
	GetUserByEmail(email string) (User, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
}

func (repo *repository) RegisterUser(user User) (User, error) {
	_, err := repo.GetUserByUID(user.UID)
	if err != nil {
		_, err = repo.db.Collection("users").InsertOne(context.Background(), user)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (repo *repository) UpdateUser(user User) (User, error) {
	userByte, err := bson.Marshal(user)
	if err != nil {
		return user, err
	}

	var User bson.M
	err = bson.Unmarshal(userByte, &User)
	if err != nil {
		return user, err
	}

	_, err = repo.db.Collection("users").UpdateOne(context.Background(), bson.M{"uid": user.UID}, bson.D{{Key: "$set", Value: User}})
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *repository) DeleteUser(id string) error {
	_, err := repo.db.Collection("users").DeleteOne(
		context.Background(),
		bson.M{
			"uid": id,
		},
	)
	if err != nil {
		return err
	}

	_, err = repo.db.Collection("activities").DeleteMany(context.Background(), bson.M{"affiliate_id": id})
	if err != nil {
		return err
	}

	_, err = repo.db.Collection("comments").DeleteMany(context.Background(), bson.M{"creator_uid": id})
	if err != nil {
		return err
	}

	_, err = repo.db.Collection("notes").DeleteMany(context.Background(), bson.M{"creator_uid": id})
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) GetUserByUID(id string) (User, error) {
	var user User

	err := repo.db.Collection("users").FindOne(
		context.Background(),
		bson.M{
			"uid": id,
		},
	).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *repository) GetUserByEmail(email string) (User, error) {
	var user User

	err := repo.db.Collection("users").FindOne(
		context.Background(),
		bson.M{
			"email": email,
		},
	).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}
