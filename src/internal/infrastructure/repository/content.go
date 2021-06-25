package repository

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/infrastructure/repository/models"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type contentRepository struct {
	db *config.DatabaseConfig
}

func NewContentRepository(db *config.DatabaseConfig) content.ContentRepository {
	return &contentRepository{db}
}

func (r *contentRepository) GetOrNone(ctx context.Context, id content.ContentID) (*content.QueryDetailContent, error) {
	c := &content.QueryDetailContent{}

	if err := models.Contents(qm.Where("id=?", id)).Bind(ctx, r.db, c); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, nil
	}

	return c, nil
}

func (r *contentRepository) GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*content.QueryContent, error) {
	favorites, err := models.Favorites(
		qm.Select("content_id"),
		qm.Where("user_id = ?", userId),
		qm.OrderBy("created_at desc"),
		qm.Limit(limit),
	).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	favoriteContentIDs := make([]int, len(favorites))
	for i, favorite := range favorites {
		favoriteContentIDs[i] = favorite.ContentID
	}

	return r.getByIDs(ctx, favoriteContentIDs)
}

func (r *contentRepository) getByIDs(ctx context.Context, ids []int) ([]*content.QueryContent, error) {
	// Ref: https://github.com/volatiletech/sqlboiler/issues/227
	convertedIDs := make([]interface{}, len(ids))
	for i, id := range ids {
		convertedIDs[i] = id
	}

	mcs, err := models.Contents(qm.WhereIn("id in ?", convertedIDs...)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	cs := make([]*content.QueryContent, len(mcs))
	for i, mc := range mcs {
		c := &content.QueryContent{
			ID:         mc.ID,
			UserID:     mc.UserID.Int,
			CategoryID: mc.CategoryID,
			Title:      mc.Title,
		}

		cs[i] = c
	}

	return cs, nil
}
