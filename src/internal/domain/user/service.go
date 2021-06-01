package user

import (
	"context"
	"io"
)

type UserService interface {
	Create(ctx context.Context, u *User) error
	Auth(ctx context.Context) (*User, error)
	Show(ctx context.Context, id string) (*User, error)
	Followings(ctx context.Context, id string) ([]*User, error)
	Followers(ctx context.Context, id string) ([]*User, error)
	UpdateThumbnail(ctx context.Context, body io.Reader) (string, error)
	Suspend(ctx context.Context, u *User) error
}
