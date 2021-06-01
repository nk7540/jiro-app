package repository

import (
	"context"

	"artics-api/src/internal/domain/follow"
	"artics-api/src/lib/models"
	"artics-api/src/lib/mysql"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type followRepository struct {
	db *mysql.Client
}

// NewFollowRepository - setups follow repository
func NewFollowRepository(db *mysql.Client) follow.FollowRepository {
	return &followRepository{
		db: db,
	}
}

func (r *followRepository) Create(ctx context.Context, f *follow.Follow) error {
	mf := models.Follow{}
	mf.ID = uuid.New().String()
	mf.FollowingID = f.FollowingID
	mf.FollowerID = f.FollowerID
	return mf.Insert(ctx, r.db.DB, boil.Infer())
}

func (r *followRepository) Delete(ctx context.Context, f *follow.Follow) error {
	mf, err := models.Follows(qm.Where("following_id=? and follower_id=?", f.FollowingID, f.FollowerID)).One(ctx, r.db.DB)
	if err != nil {
		return err
	}
	_, err = mf.Delete(ctx, r.db.DB)
	return err
}
