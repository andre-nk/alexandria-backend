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
	IsStarred        bool               `json:"is_starred" bson:"is_starred,omitempty"`
	IsArchived       bool               `json:"is_archived" bson:"is_archived,omitempty"`
	IsCommentEnabled bool               `json:"is_comment_enabled" bson:"is_comment_enabled,omitempty"`
	Collaborators    []string           `json:"collaborators,omitempty" bson:"collaborators,omitempty"`
	Views            int64              `json:"views" bson:"views,omitempty"`
}
