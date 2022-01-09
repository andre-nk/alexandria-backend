package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID               primitive.ObjectID `bson:"_id, omitempty"`
	Title            string             `bson:"title, omitempty"`
	CreatorUID       string             `bson:"creator_uid, omitempty"`
	Tags             []string           `bson:"tags, omitempty"`
	Content          []string           `bson:"content, omitempty"`
	CreatedAt        time.Time          `bson:"created_at, omitempty"`
	UpdatedAt        time.Time          `bson:"updated_at, omitempty"`
	IsStarred        bool               `bson:"is_starred, omitempty"`
	IsArchived       bool               `bson:"is_archived, omitempty"`
	IsCommentEnabled bool               `bson:"is_comment_enabled, omitempty"`
	Collaborators    []string           `bson:"collaborators, omitempty"`
	Views            int64              `bson:"views, omitempty"`
}
