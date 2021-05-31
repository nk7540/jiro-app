package service

import (
	"context"
	"strings"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/follow"
	"artics-api/src/internal/domain/user"
	"artics-api/src/middleware"

	"golang.org/x/xerrors"
)

type userService struct {
	udv user.UserDomainValidator
	ur  user.UserRepository
	fr  follow.FollowRepository
}

func NewUserService(udv user.UserDomainValidator, ur user.UserRepository, fr follow.FollowRepository) user.UserService {
	return &userService{udv, ur, fr}
}

func (s *userService) Create(ctx context.Context, u *user.User) error {
	ves := s.udv.Validate(ctx, u)
	vesPassword := s.udv.ValidatePassword(ctx, u.Password, u.PasswordConfirmation)
	ves = append(ves, vesPassword...)
	if len(ves) > 0 {
		err := xerrors.New("Failed to DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	return s.ur.Create(ctx, u)
}

func (s *userService) Auth(ctx context.Context) (*user.User, error) {
	t, err := getToken(ctx)
	if err != nil {
		return nil, err
	}

	return s.ur.GetByToken(ctx, t)
}

func (s *userService) Show(ctx context.Context, id string) (*user.User, error) {
	u, err := s.ur.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	followingCount, err := s.fr.FollowingCount(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	followerCount, err := s.fr.FollowerCount(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	u.FollowingCount = followingCount
	u.FollowerCount = followerCount

	return u, nil
}

func (s *userService) Suspend(ctx context.Context, u *user.User) error {
	return s.ur.Suspend(ctx, u)
}

func getToken(ctx context.Context) (string, error) {
	gc, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return "", xerrors.New("Cannot convert to gin.Context")
	}

	a := gc.GetHeader("Authorization")
	if a == "" {
		return "", xerrors.New("Authorization Header is not contain.")
	}

	t := strings.Replace(a, "Bearer ", "", 1)
	return t, nil
}

// OAuth認証による初回User登録時、UIDの先頭16文字を取得
// e.g.) 12345678-qwer-asdf-zxcv-uiophjklvbnm -> 12345678qwerasdf
func getName(uid string) string {
	str := strings.Replace(uid, "-", "", -1)
	return str[0:16]
}
