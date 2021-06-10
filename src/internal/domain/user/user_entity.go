package user

import (
	"artics-api/src/internal/domain/content"
	"io"
)

type User struct {
	ID               UserID
	UID              UID
	Status           Status       `validate:"required,oneof=provisional available suspended"`
	Nickname         Nickname     `validate:"max=32"`
	Email            Email        `validate:"required,email,max=256"`
	ThumbnailURL     ThumbnailURL `validate:"max=1024"`
	FavoriteContents []*content.Content
	BrowsedContents  []*content.Content
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
