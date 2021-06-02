package repository

import (
	"artics-api/src/internal/domain/content"
	"artics-api/src/lib/models"
	"artics-api/src/lib/mysql"
	"context"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type contentRepository struct {
	db *mysql.Client
}

func NewContentRepository(db *mysql.Client) content.ContentRepository {
	return &contentRepository{db}
}

func (r *contentRepository) GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*content.Content, error) {
	favoriteContents, err := models.Favorites(
		qm.Select("content_id"),
		qm.Where("user_id = ?", userId),
		qm.OrderBy("created_at desc"),
		qm.Limit(limit),
	).All(ctx, r.db.DB)
	if err != nil {
		return nil, err
	}

	favoriteContentIDs := make([]string, len(favoriteContents))
	for i, favoriteContent := range favoriteContents {
		favoriteContentIDs[i] = favoriteContent.ContentID
	}

	return r.GetByIDs(ctx, favoriteContentIDs)
}

func (r *contentRepository) GetByIDs(ctx context.Context, ids []string) ([]*content.Content, error) {
	// Ref: https://github.com/volatiletech/sqlboiler/issues/227
	convertedIDs := make([]interface{}, len(ids))
	for i, id := range ids {
		convertedIDs[i] = id
	}

	mcs, err := models.Contents(qm.WhereIn("id in ?", convertedIDs...)).All(ctx, r.db.DB)
	if err != nil {
		return nil, err
	}

	cs := make([]*content.Content, len(mcs))
	for i, mc := range mcs {
		c := &content.Content{
			UserID:     mc.UserID.String,
			CategoryID: mc.CategoryID,
			Title:      mc.Title,
		}

		cs[i] = c
	}

	return cs, nil
}
