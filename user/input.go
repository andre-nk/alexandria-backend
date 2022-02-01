package user

type UserIDUri struct {
	UID string `uri:"uid" binding:"required"`
}

type UserInput struct {
	UID         string `json:"uid" binding:"required"`
	DisplayName string `json:"displayName" binding:"required"`
	PhotoURL    string `json:"photoURL,omitempty"`
	Role        string `json:"role,omitempty"`
	Location    string `json:"location,omitempty"`
}
