package repository

import (
	"context"

	"artics-api/src/config"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/infrastructure/repository/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type followRepository struct {
	db *config.DatabaseConfig
}

// NewFollowRepository - setups follow repository
func NewFollowRepository(db *config.DatabaseConfig) user.FollowRepository {
	return &followRepository{db}
}

func (r *followRepository) Create(ctx context.Context, f *user.Follow) error {
	mf := models.Follow{
		FollowingID: int(f.FollowingID),
		FollowerID:  int(f.FollowerID),
	}
	return mf.Insert(ctx, r.db, boil.Infer())
}

func (r *followRepository) Delete(ctx context.Context, id user.FollowID) error {
	mf := &models.Follow{ID: int(id)}
	_, err := mf.Delete(ctx, r.db)
	return err
}

func (r *followRepository) GetByUserIDs(ctx context.Context, followingID user.FollowingID, followerID user.FollowerID) (*user.QueryFollow, error) {
	f := &user.QueryFollow{}
	if err := models.Follows(qm.Where("following_id=? and follower_id=?", followingID, followerID)).Bind(ctx, r.db, f); err != nil {
		return nil, err
	}
	return f, nil
}

func (r *followRepository) FollowingCount(ctx context.Context, userID user.UserID) (int, error) {
	c, err := models.Follows(qm.Where("following_id=?", int(userID))).Count(ctx, r.db)
	if err != nil {
		return 0, err
	}
	return int(c), nil
}
