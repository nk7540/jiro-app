package user

import (
	"context"

	"artics-api/src/internal/domain"
)

type UserDomainValidator interface {
	Validate(ctx context.Context, u *User) []*domain.ValidationError
	ValidatePassword(ctx context.Context, p string, pc string) []*domain.ValidationError
}
