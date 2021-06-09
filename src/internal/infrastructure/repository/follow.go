package repository

import (
	"context"

	"artics-api/src/config"
	"artics-api/src/internal/domain/follow"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/infrastructure/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type followRepository struct {
	db *config.DatabaseConfig
}

// NewFollowRepository - setups follow repository
func NewFollowRepository(db *config.DatabaseConfig) follow.FollowRepository {
	return &followRepository{db}
}

func (r *followRepository) Create(ctx context.Context, f *follow.Follow) error {
	mf := models.Follow{}
	mf.FollowingID = f.FollowingID
	mf.FollowerID = f.FollowerID
	return mf.Insert(ctx, r.db, boil.Infer())
}

func (r *followRepository) FollowingCount(ctx context.Context, userID user.ID) (int, error) {
	c, err := models.Follows(qm.Where("following_id=?", int(userID))).Count(ctx, r.db)
	if err != nil {
		return 0, err
	}
	return int(c), nil
}

func (r *followRepository) Delete(ctx context.Context, f *follow.Follow) error {
	mf, err := models.Follows(qm.Where("following_id=? and follower_id=?", f.FollowingID, f.FollowerID)).One(ctx, r.db)
	if err != nil {
		return err
	}
	_, err = mf.Delete(ctx, r.db)
	return err
}
