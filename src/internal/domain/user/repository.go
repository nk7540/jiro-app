package user

import (
	"context"
)

type UserRepository interface {
	// Command
	Create(ctx context.Context, u *User) error
	CreateWithPassword(ctx context.Context, email Email, password Password) error
	Update(ctx context.Context, cmd CommandUpdateUser) error
	Suspend(ctx context.Context, u *User) error

	// Query
	Get(ctx context.Context, id int) (*QueryDetailUser, error)
	GetByEmailOrNone(ctx context.Context, email string) (*User, error)
	GetByToken(ctx context.Context, tkn string) (*User, error)
	Followings(ctx context.Context, id int) ([]*QueryUser, error)
	Followers(ctx context.Context, id int) ([]*QueryUser, error)
}
