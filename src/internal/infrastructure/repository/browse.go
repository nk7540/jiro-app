package repository

import (
	"artics-api/src/internal/domain/browse"
	"artics-api/src/lib/models"
	"artics-api/src/lib/mysql"
	"context"
	"database/sql"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type browseRepository struct {
	db *mysql.Client
}

func NewBrowseRepository(db *mysql.Client) browse.BrowseRepository {
	return &browseRepository{db}
}

func (r *browseRepository) Save(ctx context.Context, b *browse.Browse) error {
	mb, err := models.Browses(qm.Where(
		"user_id = ? and content_id = ?",
		b.UserID,
		b.ContentID,
	)).One(ctx, r.db.DB)

	if err == sql.ErrNoRows {
		mb = &models.Browse{
			UserID:    b.UserID,
			ContentID: b.ContentID,
		}
		if err := mb.Insert(ctx, r.db.DB, boil.Infer()); err != nil {
			return err
		}
	} else if err == nil {
		mb.UpdatedAt = time.Now()
		if _, err := mb.Update(ctx, r.db.DB, boil.Whitelist("updated_at")); err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}
