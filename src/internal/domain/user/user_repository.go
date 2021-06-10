package user

import (
	"context"
)

type UserRepository interface {
	// Command
	Create(ctx context.Context, u *User) error
	CreateAuth(ctx context.Context, cmd CommandCreateUser) (UID, error)
	Update(ctx context.Context, u *User) error
	UpdateThumbnail(ctx context.Context, thumbnail Thumbnail) (ThumbnailURL, error)
	DeleteAuth(ctx context.Context, uid UID) error

	// Query
	Get(ctx context.Context, id int) (*QueryDetailUser, error)
	GetByEmailOrNone(ctx context.Context, email string) (*User, error)
	GetByToken(ctx context.Context, tkn string) (*User, error)
	Followings(ctx context.Context, id int) (*QueryUsers, error)
	Followers(ctx context.Context, id int) (*QueryUsers, error)
}
