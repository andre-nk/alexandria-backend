package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	DisplayName string             `json:"displayName" bson:"displayName"`
	Email       string             `json:"email" bson:"email"`
	PhotoURL    string             `json:"photoURL" bson:"photoURL"`
	UID         string             `json:"uid" bson:"uid"`
	Role        string             `json:"role" bson:"role"`
	Location    string             `json:"location" bson:"location"`
}
