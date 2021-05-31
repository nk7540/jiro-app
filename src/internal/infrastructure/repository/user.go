package repository

import (
	"context"
	"database/sql"

	"artics-api/src/internal/domain/user"
	"artics-api/src/lib/firebase"
	"artics-api/src/lib/models"
	"artics-api/src/lib/mysql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepository struct {
	db   *mysql.Client
	auth *firebase.Auth
}

// NewUserRepository - setups user repository
func NewUserRepository(db *mysql.Client, auth *firebase.Auth) user.UserRepository {
	return &userRepository{
		db:   db,
		auth: auth,
	}
}
func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	uid, err := r.auth.CreateUser(ctx, u.ID, u.Email, u.Password)
	if err != nil {
		return err
	}

	mu := models.User{}
	mu.ID = u.ID
	mu.UID = uid
	mu.Email = u.Email
	mu.Nickname = u.Nickname
	return mu.Insert(ctx, r.db.DB, boil.Infer())
}

func (r *userRepository) GetByToken(ctx context.Context, tkn string) (*user.User, error) {
	uid, err := r.auth.VerifyIDToken(ctx, tkn)
	if err != nil {
		return nil, err
	}

	mu, err := models.Users(qm.Where("uid = ?", uid)).One(ctx, r.db.DB)
	if err != nil {
		au, err := r.auth.GetUserByUID(ctx, uid)
		if err != nil {
			return nil, err
		}

		mu := &models.User{}
		mu.Email = au.UserInfo.Email
		mu.Insert(ctx, r.db.DB, boil.Infer())
	}
	u := &user.User{}
	u.Nickname = mu.Nickname
	u.Email = mu.Email

	return u, nil
}

func (r *userRepository) Get(ctx context.Context, id string) (*user.User, error) {
	mu, err := models.FindUser(ctx, r.db.DB, id)
	if err != nil {
		return nil, err
	}

	u := &user.User{}
	u.Nickname = mu.Nickname
	u.Email = mu.Email
	return u, err
}

func (r *userRepository) GetByEmailOrNone(ctx context.Context, email string) (*user.User, error) {
	mu, err := models.Users(qm.Where("email = ?", email)).One(ctx, r.db.DB)
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

func (r *userRepository) Followings(ctx context.Context, id string) ([]*user.User, error) {
	fs, err := models.Follows(qm.Select("follower_id"), qm.Where("following_id = ?", id)).All(ctx, r.db.DB)
	if err != nil {
		return nil, err
	}
	followingIDs := make([]string, len(fs))
	for i, f := range fs {
		followingIDs[i] = f.FollowerID
	}

	return r.getByIDs(ctx, followingIDs)
}

func (r *userRepository) Followers(ctx context.Context, id string) ([]*user.User, error) {
	fs, err := models.Follows(qm.Select("following_id"), qm.Where("follower_id = ?", id)).All(ctx, r.db.DB)
	if err != nil {
		return nil, err
	}
	followerIDs := make([]string, len(fs))
	for i, f := range fs {
		followerIDs[i] = f.FollowingID
	}

	return r.getByIDs(ctx, followerIDs)
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	mu := models.User{}
	mu.ID = u.ID
	mu.Email = u.Email
	mu.Nickname = u.Nickname
	_, err := mu.Update(ctx, r.db.DB, boil.Blacklist("uid", "status"))
	return err
}

func (r *userRepository) Suspend(ctx context.Context, u *user.User) error {
	uid, err := r.auth.GetUIDByEmail(ctx, u.Email)
	if err != nil {
		return err
	}
	if err := r.auth.DeleteUser(ctx, uid); err != nil {
		return err
	}

	mu := &models.User{}
	mu.ID = u.ID
	mu.Status = "suspended"
	mu.UID = ""

	_, err = mu.Update(ctx, r.db.DB, boil.Whitelist("status", "uid"))
	return err
}

func (r *userRepository) getByIDs(ctx context.Context, ids []string) ([]*user.User, error) {
	// Ref: https://github.com/volatiletech/sqlboiler/issues/227
	convertedIDs := make([]interface{}, len(ids))
	for i, id := range ids {
		convertedIDs[i] = id
	}

	mus, err := models.Users(qm.WhereIn("id in ?", convertedIDs...)).All(ctx, r.db.DB)
	if err != nil {
		return nil, err
	}

	us := make([]*user.User, len(mus))
	for i, mu := range mus {
		u := &user.User{
			ID:       mu.ID,
			Status:   mu.Status,
			Email:    mu.Email,
			Nickname: mu.Nickname,
		}

		us[i] = u
	}

	return us, nil
}
