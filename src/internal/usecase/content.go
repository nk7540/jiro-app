package usecase

import (
	"artics-api/src/internal/domain/content"
	"context"
)

type ContentUsecase interface {
	Show(ctx context.Context, id int) (*content.Content, error)
	Favorites(ctx context.Context, userId int) ([]*content.Content, error)
}

type contentUsecase struct {
	contentService content.ContentService
}

func NewContentUsecase(cs content.ContentService) ContentUsecase {
	return &contentUsecase{cs}
}

func (cu *contentUsecase) Show(ctx context.Context, id int) (*content.Content, error) {
	return cu.contentService.Get(ctx, id)
}

func (cu *contentUsecase) Favorites(ctx context.Context, userId int) ([]*content.Content, error) {
	return cu.contentService.GetFavoriteContents(ctx, userId, 20)
}
