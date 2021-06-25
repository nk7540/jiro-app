package user

import (
	"context"
)

type UserService interface {
	Auth(ctx context.Context, tkn string) (*User, error)
}
