package service

import (
	"context"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/browse"

	"golang.org/x/xerrors"
)

type browseService struct {
	br browse.BrowseRepository
}

func NewBrowseService(br browse.BrowseRepository) browse.BrowseService {
	return &browseService{br}
}

func (s *browseService) Save(ctx context.Context, b *browse.Browse) error {
	err := s.br.Save(ctx, b)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}
