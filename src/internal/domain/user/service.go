package user

import (
	"context"
)

type UserService interface {
	Create(ctx context.Context, u *User) error
	Auth(ctx context.Context) (*User, error)
	Show(ctx context.Context, id string) (*User, error)
	Suspend(ctx context.Context, u *User) error
}
