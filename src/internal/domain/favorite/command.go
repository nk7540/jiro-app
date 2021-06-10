package favorite

type CommandLike struct {
	UserID    UserID
	ContentID ContentID
}

type CommandUnlike struct {
	UserID    UserID
	ContentID ContentID
}
