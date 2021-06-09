package user

type CommandUpdateUser struct {
	ID           int
	Nickname     string `validate:"max=32"`
	ThumbnailURL string `validate:"max=256"`
}
