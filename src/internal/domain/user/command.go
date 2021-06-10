package user

type CommandCreateUser struct {
	Email                string `validate:"required,email,max=256"`
	Password             string `validate:"required,password"`
	PasswordConfirmation string `validate:"required,equalTo=password"`
}

type CommandUpdateUser struct {
	Nickname     string `validate:"max=32"`
	ThumbnailURL string `validate:"max=256"`
}
