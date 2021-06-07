package user

import (
	"context"
	"io"
)

type UserService interface {
	Create(ctx context.Context, u *User) error
	Auth(ctx context.Context, tkn string) (*User, error)
	Show(ctx context.Context, id int) (*User, error)
	Followings(ctx context.Context, id int) ([]*User, error)
	Followers(ctx context.Context, id int) ([]*User, error)
	UpdateThumbnail(ctx context.Context, body io.Reader) (string, error)
	Update(ctx context.Context, u *User) error
	Suspend(ctx context.Context, u *User) error
}
