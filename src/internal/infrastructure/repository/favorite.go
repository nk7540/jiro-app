package repository

import (
	"context"
	"database/sql"

	"artics-api/src/config"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/infrastructure/repository/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type favoriteRepository struct {
	db *config.DatabaseConfig
}

func NewFavoriteRepository(db *config.DatabaseConfig) content.FavoriteRepository {
	return &favoriteRepository{db}
}

func (r *favoriteRepository) Create(ctx context.Context, f *content.Favorite) error {
	mf := models.Favorite{
		UserID:    int(f.UserID),
		ContentID: int(f.ContentID),
	}
	return mf.Insert(ctx, r.db, boil.Infer())
}

func (r *favoriteRepository) Delete(ctx context.Context, id content.FavoriteID) error {
	mf := models.Favorite{ID: int(id)}
	_, err := mf.Delete(ctx, r.db)
	return err
}

func (r *favoriteRepository) FindByUserAndContentIDOrNone(ctx context.Context, userID content.FavoriteUserID, contentID content.FavoriteContentID) (*content.QueryFavorite, error) {
	f := &content.QueryFavorite{}
	if err := models.Favorites(qm.Where(
		"user_id = ? and content_id = ?",
		f.UserID,
		f.ContentID,
	)).Bind(ctx, r.db, f); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return f, nil
}
