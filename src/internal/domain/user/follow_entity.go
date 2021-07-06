package user

// Follow entity
type Follow struct {
	ID          FollowID
	FollowingID FollowingID
	FollowerID  FollowerID
	IsClose     FollowIsClose
}

type FollowID int
type FollowingID int
type FollowerID int
type FollowIsClose bool
