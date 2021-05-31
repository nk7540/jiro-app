package usecase

import (
	"artics-api/src/internal/domain/content"
	"context"
)

type ContentUsecase interface {
	Favorites(ctx context.Context, userId string) ([]*content.Content, error)
}

type contentUsecase struct {
	contentService content.ContentService
}

func NewContentUsecase(cs content.ContentService) ContentUsecase {
	return &contentUsecase{cs}
}

func (cu *contentUsecase) Favorites(ctx context.Context, userId string) ([]*content.Content, error) {
	return cu.contentService.GetFavoriteContents(ctx, userId, 20)
}
