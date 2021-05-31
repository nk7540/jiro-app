package service

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/favorite"
	"context"

	"golang.org/x/xerrors"
)

type favoriteService struct {
	fr favorite.FavoriteRepository
}

func NewFavoriteService(fr favorite.FavoriteRepository) favorite.FavoriteService {
	return &favoriteService{fr}
}

func (s *favoriteService) Create(ctx context.Context, f *favorite.Favorite) error {
	err := s.fr.Create(ctx, f)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (s *favoriteService) Delete(ctx context.Context, f *favorite.Favorite) error {
	err := s.fr.Delete(ctx, f)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}
