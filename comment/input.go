package comment

type CommentIDUri struct {
	ID string `uri:"comment_id" binding:"required"`
}

type CreateCommentInput struct {
	NoteID     string   `json:"note_id,omitempty"`
	CreatorUID string   `json:"creator_uid,omitempty"`
	Content    string   `json:"content,omitempty"`
	Mentions   []string `json:"mentions,omitempty"`
}
