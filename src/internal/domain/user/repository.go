package user

import (
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, u *User) error
	Show(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, u *User) error
}
