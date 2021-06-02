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

func (r *contentRepository) Get(ctx context.Context, id int) (*content.Content, error) {
	mc, err := models.FindContent(ctx, r.db.DB, id)
	if err != nil {
		return nil, err
	}

	c := &content.Content{
		UserID:     mc.UserID.Int,
		CategoryID: mc.CategoryID,
		Title:      mc.Title,
	}

	return c, nil
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

	favoriteContentIDs := make([]int, len(favoriteContents))
	for i, favoriteContent := range favoriteContents {
		favoriteContentIDs[i] = favoriteContent.ContentID
	}

	return r.getByIDs(ctx, favoriteContentIDs)
}

func (r *contentRepository) getByIDs(ctx context.Context, ids []int) ([]*content.Content, error) {
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
			UserID:     mc.UserID.Int,
			CategoryID: mc.CategoryID,
			Title:      mc.Title,
		}

		cs[i] = c
	}

	return cs, nil
}
