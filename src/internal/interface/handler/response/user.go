package response

type CreateUser struct {
	ResultCode string `json:"resultCode"`
}

type ShowUser struct {
	ID             int    `json:"id"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Followingcount int    `json:"followingCount"`
	Followercount  int    `json:"followerCount"`
}

type Users struct {
	Users []*User `json:"users"`
}

type User struct {
	ID           int    `json:"id"`
	Nickname     string `json:"nickname"`
	ThumbnailURL string `json:"thumbnailURL"`
}

type UpdateUser struct {
	ID           int    `json:"id"`
	Nickname     string `json:"nickname"`
	ThumbnailURL string `json:"thumbnailURL"`
}
