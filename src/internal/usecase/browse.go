package usecase

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/browse"
	"artics-api/src/internal/domain/user"
	"context"
)

type BrowseUsecase interface {
	Save(ctx context.Context, contentID string) error
}

type browseUsecase struct {
	userService   user.UserService
	browseService browse.BrowseService
}

func NewBrowseUsecase(us user.UserService, bs browse.BrowseService) BrowseUsecase {
	return &browseUsecase{us, bs}
}

func (bu *browseUsecase) Save(ctx context.Context, contentID string) error {
	u, err := bu.userService.Auth(ctx)
	if err != nil {
		return domain.Unauthorized.New(err)
	}

	b := &browse.Browse{
		UserID:    u.ID,
		ContentID: contentID,
	}

	return bu.browseService.Save(ctx, b)
}
