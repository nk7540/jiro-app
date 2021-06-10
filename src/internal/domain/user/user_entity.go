package user

import (
	"fmt"
	"io"
)

type User struct {
	ID           UserID
	UID          UID
	Status       Status       `validate:"required,oneof=provisional available suspended"`
	Nickname     Nickname     `validate:"max=32"`
	Email        Email        `validate:"required,email,max=256"`
	ThumbnailURL ThumbnailURL `validate:"max=1024"`
}

type UserID int
type UID string
type Status string
type Nickname string
type Email string
type ThumbnailURL string

// Virtual attributes
type Password string
type PasswordConfirmation string
type Thumbnail io.Reader

const (
	Provisional = Status("provisional")
	Available   = Status("available")
	Suspended   = Status("suspended")
)

func (u *User) Suspend() {
	u.Status = Suspended
	u.Nickname = ""
	u.Email = Email(fmt.Sprintf("leaved+user%s@artics.jp", u.ID))
	u.ThumbnailURL = ""
}
