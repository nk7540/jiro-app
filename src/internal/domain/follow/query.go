package follow

type QueryFollow struct {
	ID          int `boil:"id"`
	FollowingID int `boil:"following_id"`
	FollowerID  int `boil:"follower_id"`
}
