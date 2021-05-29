package user

import (
	"context"
)

type UserService interface {
	CreateAuth(ctx context.Context, u *User) (string, error)
	Auth(ctx context.Context) (string, error)
	DeleteAuth(ctx context.Context, u *User) error
}
