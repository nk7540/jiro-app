package repository

import (
	"context"
	"database/sql"

	"artics-api/src/config"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/infrastructure/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepository struct {
	db   *config.DatabaseConfig
	auth *config.AuthConfig
}

// NewUserRepository - setups user repository
func NewUserRepository(db *config.DatabaseConfig, auth *config.AuthConfig) user.UserRepository {
	return &userRepository{db, auth}
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	mu := models.User{
		UID:   string(u.UID),
		Email: string(u.Email),
	}
	return mu.Insert(ctx, r.db, boil.Infer())
}

func (r *userRepository) CreateAuth(ctx context.Context, cmd user.CommandCreateUser) (user.UID, error) {
	uid, err := r.auth.CreateUser(ctx, cmd.Email, cmd.Password)
	if err != nil {
		return "", err
	}
	return user.UID(uid), nil
}

func (r *userRepository) GetByToken(ctx context.Context, tkn string) (*user.User, error) {
	uid, err := r.auth.VerifyIDToken(ctx, tkn)
	if err != nil {
		return nil, err
	}

	mu, err := models.Users(qm.Where("uid = ?", uid)).One(ctx, r.db)
	if err != nil {
		au, err := r.auth.GetUserByUID(ctx, uid)
		if err != nil {
			return nil, err
		}

		mu := &models.User{}
		mu.Email = au.UserInfo.Email
		mu.Insert(ctx, r.db, boil.Infer())
	}
	u := &user.User{}
	u.Nickname = mu.Nickname
	u.Email = mu.Email

	return u, nil
}

func (r *userRepository) Get(ctx context.Context, id int) (*user.QueryDetailUser, error) {
	u := &user.QueryDetailUser{}
	if err := models.Users(qm.Where("id=?", id)).Bind(ctx, r.db, u); err != nil {
		return nil, err
	}

	followingCount, err := models.Follows(qm.Where("following_id=?", id)).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}
	followerCount, err := models.Follows(qm.Where("follower_id=?", id)).Count(ctx, r.db)
	if err != nil {
		return nil, err
	}
	u.FollowingCount = int(followingCount)
	u.FollowerCount = int(followerCount)

	return u, nil
}

func (r *userRepository) GetByEmailOrNone(ctx context.Context, email string) (*user.User, error) {
	mu, err := models.Users(qm.Where("email = ?", email)).One(ctx, r.db)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user.User{
		ID:       mu.ID,
		Status:   mu.Status,
		Email:    mu.Email,
		Nickname: mu.Nickname,
	}, nil
}

func (r *userRepository) Followings(ctx context.Context, id int) (*user.QueryUsers, error) {
	fs, err := models.Follows(qm.Select("follower_id"), qm.Where("following_id = ?", id)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	followingIDs := make([]int, len(fs))
	for i, f := range fs {
		followingIDs[i] = f.FollowerID
	}

	return r.getByIDs(ctx, followingIDs)
}

func (r *userRepository) Followers(ctx context.Context, id int) (*user.QueryUsers, error) {
	fs, err := models.Follows(qm.Select("following_id"), qm.Where("follower_id = ?", id)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	followerIDs := make([]int, len(fs))
	for i, f := range fs {
		followerIDs[i] = f.FollowingID
	}

	return r.getByIDs(ctx, followerIDs)
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	mu := models.User{
		ID:           int(u.ID),
		Status:       string(u.Status),
		Nickname:     string(u.Nickname),
		ThumbnailURL: string(u.ThumbnailURL),
	}
	_, err := mu.Update(ctx, r.db, boil.Blacklist("uid"))
	return err
}

func (r *userRepository) DeleteAuth(ctx context.Context, uid user.UID) error {
	return r.auth.DeleteUser(ctx, string(uid))
}

func (r *userRepository) getByIDs(ctx context.Context, ids []int) (*user.QueryUsers, error) {
	// Ref: https://github.com/volatiletech/sqlboiler/issues/227
	convertedIDs := make([]interface{}, len(ids))
	for i, id := range ids {
		convertedIDs[i] = id
	}

	us := new(user.QueryUsers)
	if err := models.Users(qm.WhereIn("id in ?", convertedIDs...)).Bind(ctx, r.db, us); err != nil {
		return nil, err
	}

	return us, nil
}
