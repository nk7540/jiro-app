package content

type CommandLike struct {
	UserID    FavoriteUserID
	ContentID FavoriteContentID
}

type CommandUnlike struct {
	UserID    FavoriteUserID
	ContentID FavoriteContentID
}
