package usecase

import (
	"artics-api/src/internal/domain/browse"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
	"context"
)

type BrowseUsecase interface {
	Save(ctx context.Context, contentID int) error
}

type browseUsecase struct {
	userService   user.UserService
	browseService browse.BrowseService
}

func NewBrowseUsecase(us user.UserService, bs browse.BrowseService) BrowseUsecase {
	return &browseUsecase{us, bs}
}

func (bu *browseUsecase) Save(ctx context.Context, contentID int) error {
	c := ctx.(pkg.Context)
	u := c.Locals("user").(user.User)

	b := &browse.Browse{
		UserID:    u.ID,
		ContentID: contentID,
	}

	return bu.browseService.Save(ctx, b)
}
