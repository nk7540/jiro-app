package user

type QueryDetailUser struct {
	ID             int    `boil:"id"`
	Nickname       string `boil:"nickname"`
	ThumbnailURL   string `boil:"thumbnail_url"`
	Profile        string `boil:"profile"`
	FollowingCount int    `boil:"-"`
	FollowerCount  int    `boil:"-"`
}

type QueryUsers struct {
	Users []*QueryUser
}

type QueryUser struct {
	ID           int    `boil:"id"`
	Nickname     string `boil:"nickname"`
	ThumbnailURL string `boil:"thumbnail_url"`
}
