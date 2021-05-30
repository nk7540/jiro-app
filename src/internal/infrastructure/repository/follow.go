package repository

import (
	"context"

	"artics-api/src/internal/domain/follow"
	"artics-api/src/lib/mysql"
	"artics-api/src/lib/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/google/uuid"
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

func (r *followRepository) FollowingCount(ctx context.Context, id string) (int, error) {
	cnt, err := models.Follows(qm.Where("following_id=?", id)).Count(ctx, r.db.DB)
	return int(cnt), err
}

func (r *followRepository) FollowerCount(ctx context.Context, id string) (int, error) {
	cnt, err := models.Follows(qm.Where("follower_id=?", id)).Count(ctx, r.db.DB)
	return int(cnt), err
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
