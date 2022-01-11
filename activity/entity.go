package activity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activity struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ActivityID  primitive.ObjectID `json:"activity_id,omitempty" bson:"activity_id,omitempty"`
	AffiliateID string             `json:"affiliate_id,omitempty" bson:"affiliate_id,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	IsRead      bool               `json:"is_read" bson:"is_read"`
	Message     string             `json:"message,omitempty" bson:"message,omitempty"`
}
