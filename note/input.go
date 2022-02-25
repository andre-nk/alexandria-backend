package note

type NoteIDUri struct {
	ID string `uri:"id" binding:"required"`
}

type CreateNoteInput struct {
	Title                string   `json:"title" binding:"required"`
	CreatorUID           string   `json:"creator_uid" binding:"required"`
	Tags                 []string `json:"tags" binding:"required"`
	Content              string   `json:"content" binding:"required"`
	IsStarred            bool     `json:"is_starred"`
	IsCommentEnabled     bool     `json:"is_comment_enabled" binding:"required"`
	PendingCollaborators []string `json:"pending_collaborators" binding:"required"`
	Collaborators        []string `json:"collaborators" binding:"required"`
}

type UpdateNoteInput struct {
	ID                  NoteIDUri `json:"_id" binding:"required"`
	CreatorUID          string    `json:"creator_uid" binding:"required"`
	Title               string    `json:"title" binding:"required"`
	Tags                []string  `json:"tags" binding:"required"`
	Content             string    `json:"content" binding:"required"`
	IsStarred           bool      `json:"is_starred"`
	IsArchived          bool      `json:"is_archived"`
	IsCommentEnabled    bool      `json:"is_comment_enabled"`
	PendingCollaboratos []string  `json:"pending_collaborators"`
	Collaborators       []string  `json:"collaborators" binding:"required"`
	Views               int64     `json:"views"`
}
