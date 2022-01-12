package user

type UserIDUri struct {
	UID string `uri:"uid" binding:"required"`
}

type UserInput struct {
	UID      string   `json:"uid" binding:"required"`
	Role     string   `json:"role,omitempty"`
	Location string   `json:"location,omitempty"`
	Friends  []string `json:"friends,omitempty" binding:"required"`
}
