package user

type QueryDetailUser struct {
	ID             int
	Nickname       string
	ThumbnailURL   string
	Profile        string
	FollowingCount int
	FollowerCount  int
}
