package repository

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/infrastructure/repository/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type noticeRepository struct {
	db *config.DatabaseConfig
}

func NewNoticeRepository(db *config.DatabaseConfig) content.NoticeRepository {
	return &noticeRepository{db}
}

func (r *noticeRepository) Create(ctx context.Context, n *content.Notice) error {
	mn := models.Notice{
		UserID: int(n.UserID),
		Title:  string(n.Title),
		Body:   string(n.Body),
	}
	return mn.Insert(ctx, r.db, boil.Infer())
}
