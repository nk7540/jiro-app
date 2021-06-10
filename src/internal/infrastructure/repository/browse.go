package repository

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/infrastructure/models"
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type browseRepository struct {
	db *config.DatabaseConfig
}

func NewBrowseRepository(db *config.DatabaseConfig) content.BrowseRepository {
	return &browseRepository{db}
}

func (r *browseRepository) Save(ctx context.Context, b *content.Browse) error {
	mb, err := models.Browses(qm.Where(
		"user_id = ? and content_id = ?",
		b.UserID,
		b.ContentID,
	)).One(ctx, r.db)

	if err == sql.ErrNoRows {
		mb = &models.Browse{
			UserID:    int(b.UserID),
			ContentID: int(b.ContentID),
		}
		if err := mb.Insert(ctx, r.db, boil.Infer()); err != nil {
			return err
		}
	} else if err == nil {
		mb.UpdatedAt = time.Now()
		if _, err := mb.Update(ctx, r.db, boil.Whitelist("updated_at")); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}
