package service

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"context"

	"golang.org/x/xerrors"
)

type contentService struct {
	cr content.ContentRepository
}

func NewContentService(cr content.ContentRepository) content.ContentService {
	return &contentService{cr}
}

func (s *contentService) Get(ctx context.Context, id int) (*content.Content, error) {
	c, err := s.cr.Get(ctx, id)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return nil, domain.NotFound.New(err)
	}

	return c, nil
}

func (s *contentService) GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*content.Content, error) {
	cs, err := s.cr.GetFavoriteContents(ctx, userId, limit)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return cs, nil
}
