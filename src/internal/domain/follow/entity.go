package follow

// Follow entity
type Follow struct {
	ID          ID
	FollowingID FollowingID
	FollowerID  FollowerID
}

type ID int
type FollowingID int
type FollowerID int
