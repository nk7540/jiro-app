package user

type CommandCreateUser struct {
	Email                string `validate:"required,email,max=256"`
	Password             string `validate:"required,password"`
	PasswordConfirmation string `validate:"required,equalTo=password"`
}

type CommandUpdateUser struct {
	Nickname     Nickname     `validate:"max=32"`
	ThumbnailURL ThumbnailURL `validate:"max=256"`
}

type CommandFollow struct {
	User       *User
	FollowerID FollowerID
}

type CommandUnfollow struct {
	FollowingID FollowingID
	FollowerID  FollowerID
}
