package user

type QueryDetailUser struct {
	ID             int
	Nickname       string
	ThumbnailURL   string
	Profile        string
	FollowingCount int
	FollowerCount  int
}

type QueryUser struct {
	ID           int
	Nickname     string
	ThumbnailURL string
}
