package comment

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentIDUri struct {
	ID string `uri:"id" binding:"required"`
}

type CreateCommentInput struct {
	NoteID     primitive.ObjectID `json:"note_id,omitempty"`
	CreatorUID string             `json:"creator_uid,omitempty"`
	Content    string             `json:"content,omitempty"`
	Mentions   []string           `json:"mentions,omitempty"`
}
