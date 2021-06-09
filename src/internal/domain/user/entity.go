package user

import (
	"artics-api/src/internal/domain/content"
)

type User struct {
	ID               ID
	Status           Status   `validate:"required,oneof=provisional available suspended"`
	Nickname         Nickname `validate:"max=256"`
	Email            Email    `validate:"required,email,max=256"`
	ThumbnailURL     ThumbnailURL
	FavoriteContents []*content.Content
	BrowsedContents  []*content.Content
}

type ID int
type Status string
type Nickname string
type Email string
type ThumbnailURL string

// Virtual attributes
type UID string
type Password string
type PasswordConfirmation string
