package follow

type CommandFollow struct {
	FollowingID FollowingID
	FollowerID  FollowerID
}

type CommandUnfollow struct {
	FollowingID FollowingID
	FollowerID  FollowerID
}
