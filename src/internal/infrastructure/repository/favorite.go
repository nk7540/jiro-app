package repository

import (
	"context"

	"artics-api/src/internal/domain/favorite"
	"artics-api/src/lib/models"
	"artics-api/src/lib/mysql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type favoriteRepository struct {
	db *mysql.Client
}

func NewFavoriteRepository(db *mysql.Client) favorite.FavoriteRepository {
	return &favoriteRepository{db}
}

func (r *favoriteRepository) Create(ctx context.Context, f *favorite.Favorite) error {
	mf := models.Favorite{
		ID:        f.ID,
		UserID:    f.UserID,
		ContentID: f.ContentID,
	}
	return mf.Insert(ctx, r.db.DB, boil.Infer())
}
