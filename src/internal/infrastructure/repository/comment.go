package repository

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/infrastructure/repository/models"
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type commentRepository struct {
	db *config.DatabaseConfig
}

func NewCommentRepository(db *config.DatabaseConfig) content.CommentRepository {
	return &commentRepository{db}
}

func (r *commentRepository) Create(ctx context.Context, c *content.Comment) (content.CommentID, error) {
	mc := models.Comment{
		UserID:    int(c.UserID),
		ContentID: int(c.ContentID),
		Body:      string(c.Body),
	}
	if err := mc.Insert(ctx, r.db, boil.Infer()); err != nil {
		return 0, err
	}
	return content.CommentID(mc.ID), nil
}
