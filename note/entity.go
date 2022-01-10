package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title            string             `json:"title,omitempty" bson:"title,omitempty"`
	CreatorUID       string             `json:"creator_uid,omitempty" bson:"creator_uid,omitempty"`
	Tags             []string           `json:"tags,omitempty" bson:"tags,omitempty"`
	Content          []string           `json:"content,omitempty" bson:"content,omitempty"`
	CreatedAt        time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt        time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	IsStarred        bool               `json:"is_starred,omitempty" bson:"is_starred"`
	IsArchived       bool               `json:"is_archived,omitempty" bson:"is_archived"`
	IsCommentEnabled bool               `json:"is_comment_enabled,omitempty" bson:"is_comment_enabled"`
	Collaborators    []string           `json:"collaborators,omitempty" bson:"collaborators,omitempty"`
	Views            int64              `json:"views,omitempty" bson:"views,omitempty"`
}
