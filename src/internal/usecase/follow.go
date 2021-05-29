package usecase

import (
	"context"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/follow"
)

type FollowUsecase interface {
	Create(ctx context.Context, id int) error
	Destroy(ctx context.Context, id int) error
}

type followUsecase struct {
	followRepository follow.FollowRepository
}

// NewFollowUsecase - generates follow usecase
func NewFollowUsecase(fr follow.FollowRepository) *FollowUsecase {
	return &followUsecase{fr}
}

func (fu *followUsecase) Create(ctx context.Context, id int) error {
	u, err := uu.userService.Auth(ctx)
	if err != nil {
		return nil, domain.Unauthorized.New(err)
	}

	f := &follow.Follow{}
	f.FollowingID = u.ID
	f.FollowerID = id
	if err := uu.followRepository.Create(ctx, f); err != nil {
		return nil, err
	}

	return nil
}

func (fu *followUsecase) Destroy(ctx context.Context, id int) error {
	u, err := uu.userService.Auth(ctx)
	if err != nil {
		return nil, domain.Unauthorized.New(err)
	}

	f := &follow.Follow{}
	f.FollowingID = u.ID
	f.FollowerID = id
	if err := uu.followRepository.Destroy(ctx, f); err != nil {
		return nil, err
	}

	return nil
}
