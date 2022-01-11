package user

type UserIDUri struct {
	UID string `uri:"uid" binding:"required"`
}
