package usecase

import (
	"artics-api/src/internal/domain/favorite"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
	"context"
)

type FavoriteUsecase interface {
	Create(ctx context.Context, contentID int) error
	Delete(ctx context.Context, contentID int) error
}

type favoriteUsecase struct {
	userService     user.UserService
	favoriteService favorite.FavoriteService
}

func NewFavoriteUsecase(us user.UserService, fs favorite.FavoriteService) FavoriteUsecase {
	return &favoriteUsecase{us, fs}
}

func (fu *favoriteUsecase) Create(ctx context.Context, contentID int) error {
	c := ctx.(pkg.Context)
	u := c.Locals("user").(user.User)

	f := &favorite.Favorite{
		UserID:    u.ID,
		ContentID: contentID,
	}

	return fu.favoriteService.Create(ctx, f)
}

func (fu *favoriteUsecase) Delete(ctx context.Context, contentID int) error {
	c := ctx.(pkg.Context)
	u := c.Locals("user").(user.User)

	f := &favorite.Favorite{
		UserID:    u.ID,
		ContentID: contentID,
	}

	return fu.favoriteService.Delete(ctx, f)
}
