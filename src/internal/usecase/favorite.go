package usecase

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/favorite"
	"artics-api/src/internal/domain/user"
	"context"

	"github.com/google/uuid"
)

type FavoriteUsecase interface {
	Create(ctx context.Context, contentID string) error
	Delete(ctx context.Context, contentID string) error
}

type favoriteUsecase struct {
	userService     user.UserService
	favoriteService favorite.FavoriteService
}

func NewFavoriteUsecase(us user.UserService, fs favorite.FavoriteService) FavoriteUsecase {
	return &favoriteUsecase{us, fs}
}

func (fu *favoriteUsecase) Create(ctx context.Context, contentID string) error {
	u, err := fu.userService.Auth(ctx)
	if err != nil {
		return domain.Unauthorized.New(err)
	}

	f := &favorite.Favorite{
		ID:        uuid.New().String(),
		UserID:    u.ID,
		ContentID: contentID,
	}

	return fu.favoriteService.Create(ctx, f)
}

func (fu *favoriteUsecase) Delete(ctx context.Context, contentID string) error {
	u, err := fu.userService.Auth(ctx)
	if err != nil {
		return domain.Unauthorized.New(err)
	}

	f := &favorite.Favorite{
		UserID:    u.ID,
		ContentID: contentID,
	}

	return fu.favoriteService.Delete(ctx, f)
}
