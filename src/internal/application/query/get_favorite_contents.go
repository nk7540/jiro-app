package query

import (
	"artics-api/src/internal/domain/content"
	"context"
)

type GetFavoriteContentsHanlder struct {
	contentRepository content.ContentRepository
}

func NewGetFavoriteContentsHandler(cr content.ContentRepository) GetFavoriteContentsHanlder {
	return GetFavoriteContentsHanlder{cr}
}

type GetFavoriteContents struct {
	UserID int
	Limit  int
}

func (h GetFavoriteContentsHanlder) Handle(ctx context.Context, query GetFavoriteContents) ([]*content.QueryContent, error) {
	return h.contentRepository.GetFavoriteContents(ctx, query.UserID, query.Limit)
}
