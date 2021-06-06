package usecase

import (
	"context"

	"artics-api/src/internal/domain/follow"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type FollowUsecase interface {
	Create(ctx context.Context, id int) error
	Delete(ctx context.Context, id int) error
}

type followUsecase struct {
	followRepository follow.FollowRepository
	userService      user.UserService
}

// NewFollowUsecase - generates follow usecase
func NewFollowUsecase(fr follow.FollowRepository, us user.UserService) FollowUsecase {
	return &followUsecase{fr, us}
}

func (fu *followUsecase) Create(ctx context.Context, id int) error {
	c := ctx.(pkg.Context)
	u := c.Locals("user").(user.User)

	f := &follow.Follow{}
	f.FollowingID = u.ID
	f.FollowerID = id

	return fu.followRepository.Create(ctx, f)
}

func (fu *followUsecase) Delete(ctx context.Context, id int) error {
	c := ctx.(pkg.Context)
	u := c.Locals("user").(user.User)

	f := &follow.Follow{}
	f.FollowingID = u.ID
	f.FollowerID = id
	return fu.followRepository.Delete(ctx, f)
}
