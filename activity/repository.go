package activity

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateActivity(activity Activity) (Activity, error)
	GetActivityByAffiliateID(id string) ([]Activity, error)
}

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *repository {
	return &repository{db}
}

func (repo *repository) CreateActivity(activity Activity) (Activity, error) {
	_, err := repo.db.Collection("activities").InsertOne(context.Background(), activity)
	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (repo *repository) GetActivityByAffiliateID(id string) ([]Activity, error) {
	affiliateID, _ := primitive.ObjectIDFromHex(id)

	cursor, err := repo.db.Collection("activities").Find(context.Background(), bson.M{"affiliate_id": affiliateID})
	if err != nil {
		return []Activity{}, nil
	}

	var activities []Activity
	err = cursor.All(context.Background(), &activities)
	if err != nil {
		return activities, err
	}

	return activities, nil
}
