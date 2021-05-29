package repository

import (
	"context"

	"artics-api/src/internal/domain/models"
	"artics-api/src/internal/domain/user"
	"artics-api/src/lib/mysql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type userRepository struct {
	db *mysql.Client
}

// NewUserRepository - setups user repository
func NewUserRepository(db *mysql.Client) user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	return u.Insert(ctx, r.db.DB, boil.Infer())
}

func (r *userRepository) Show(ctx context.Context, id int) (*user.User, error) {
	u, err := models.FindUser(ctx, r.db.DB, id)
	if err != nil {
		return nil, err
	}

	return &user.User{*u, ""}, err
}

func (r *userRepository) Update(ctx context.Context, u *user.User) error {
	_, err := u.Update(ctx, r.db.DB, boil.Infer())
	return err;
}
