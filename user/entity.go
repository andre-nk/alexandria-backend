package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UID      string             `json:"uid" bson:"uid"`
	Role     string             `json:"role" bson:"role"`
	Location string             `json:"location" bson:"location"`
	Friends  []string           `json:"friends" bson:"friends"`
}
