package note

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`                 //biarin gak mungkin gak diisi
	Title            string             `json:"title,omitempty" bson:"title,omitempty"`             //biarin nanti di formik ini required
	CreatorUID       string             `json:"creator_uid,omitempty" bson:"creator_uid,omitempty"` //biarin pasti ada
	Tags             []string           `json:"tags,omitempty" bson:"tags"`                         //no omitempty spy bisa dihapus tagsnya / no tags at all
	Content          string             `json:"content,omitempty" bson:"content"`                   //no omitempty spy bisa dihapus
	CreatedAt        time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`   //biarin
	UpdatedAt        time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`   //biarin
	IsStarred        bool               `json:"is_starred" bson:"is_starred"`                       //no omitempty spy bisa false
	IsArchived       bool               `json:"is_archived" bson:"is_archived"`                     //no omitempty spy bisa false
	IsCommentEnabled bool               `json:"is_comment_enabled" bson:"is_comment_enabled"`       //no omitempty spy bisa false
	Collaborators    []string           `json:"collaborators" bson:"collaborators"`                 //no omitempty spy bisa []
	Views            int64              `json:"views" bson:"views,omitempty"`                       //biarin karena view pertama => 1
}
