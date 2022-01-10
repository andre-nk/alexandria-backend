package comment

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NoteID     primitive.ObjectID `json:"note_id,omitempty" bson:"note_id,omitempty"`
	CreatorUID string             `json:"creator_uid,omitempty" bson:"creator_uid,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	Content    string             `json:"content,omitempty" bson:"content,omitempty"`
	Mentions   []string           `json:"mentions,omitempty" bson:"mentions,omitempty"`
}
