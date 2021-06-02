package user

import (
	"artics-api/src/internal/domain/content"
)

// User entity
type User struct {
	ID                   int
	Status               string `validate:"required,oneof=provisional available suspended"`
	Nickname             string `validate:"max=256"`
	Email                string `validate:"required,email,max=256"`
	ThumbnailURL         string
	Password             string
	PasswordConfirmation string
	FollowingCount       int
	FollowerCount        int
	FavoriteContents     []*content.Content
	BrowsedContents      []*content.Content
}
