package repository

import (
	"context"
	"database/sql"

	"artics-api/src/internal/domain/favorite"
	"artics-api/src/lib/models"
	"artics-api/src/lib/mysql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type favoriteRepository struct {
	db *mysql.Client
}

func NewFavoriteRepository(db *mysql.Client) favorite.FavoriteRepository {
	return &favoriteRepository{db}
}

func (r *favoriteRepository) Create(ctx context.Context, f *favorite.Favorite) error {
	mf := models.Favorite{
		UserID:    f.UserID,
		ContentID: f.ContentID,
	}
	return mf.Insert(ctx, r.db.DB, boil.Infer())
}

func (r *favoriteRepository) Delete(ctx context.Context, f *favorite.Favorite) error {
	mf, err := models.Favorites(qm.Where(
		"user_id = ? and content_id = ?",
		f.UserID,
		f.ContentID,
	)).One(ctx, r.db.DB)

	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		return err
	}

	_, err = mf.Delete(ctx, r.db.DB)
	return err
}
