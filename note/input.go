package note

type NoteIDUri struct {
	ID string `uri:"id" binding:"required"`
}

type CreateNoteInput struct {
	Title            string   `json:"title" binding:"required"`
	CreatorUID       string   `json:"creator_uid" binding:"required"`
	Tags             []string `json:"tags" binding:"required"`
	Content          []string `json:"content" binding:"required"`
	IsStarred        bool     `json:"is_starred" binding:"required"`
	IsCommentEnabled bool     `json:"is_comment_enabled" binding:"required"`
	Collaborators    []string `json:"collaborators" binding:"required"`
}

type UpdateNoteInput struct {
	ID               NoteIDUri `json:"_id" binding:"required"`
	CreatorUID       string    `json:"creator_uid"`
	Title            string    `json:"title"`
	Tags             []string  `json:"tags"`
	Content          []string  `json:"content"`
	IsStarred        bool      `json:"is_starred"`
	IsArchived       bool      `json:"is_archived"`
	IsCommentEnabled bool      `json:"is_comment_enabled"`
	Collaborators    []string  `json:"collaborators"`
	Views            int64     `json:"views"`
}
