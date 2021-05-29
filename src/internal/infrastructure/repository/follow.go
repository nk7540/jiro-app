package repository

import (
	"context"

	"artics-api/src/internal/domain/models"
	"artics-api/src/internal/domain/follow"
	"artics-api/src/lib/mysql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type followRepository struct {
odb *mysql.Client
}

// NewFollowRepository - setups follow repository
func NewFollowRepository(db *mysql.Client) follow.FollowRepository {
	return &followRepository{
		db: db,
	}
}

func (r *followRepository) FollowingCount(ctx context.Context, id int) (int, error) {
	return follow.Follow(Where("following_id=?", id)).Count(ctx, r.db.DB)
}

func (r *followRepository) FollowerCount(ctx context.Context, id int) (int, error) {
	return follow.Follow(Where("follower_id=?", id)).Count(ctx, r.db.DB)
}

func (r *followRepository) Create(ctx context.Context, f *follow.Follow) error {
	_, err := f.Insert(ctx, r.db.DB, boil.Infer()))
	return err
}

func (r *followRepository) Delete(ctx context.Context, follow *follow.Follow) error {
	_, err := f.Delete(ctx, r.db.DB)
	return err
}
