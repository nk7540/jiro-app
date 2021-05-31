package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	Get(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByToken(ctx context.Context, tkn string) (*User, error)
	Update(ctx context.Context, u *User) error
	Suspend(ctx context.Context, u *User) error
}
