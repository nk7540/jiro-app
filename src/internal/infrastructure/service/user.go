package service

import (
	"context"
	"strings"

	"artics-api/src/internal/domain/user"
	"artics-api/src/lib/firebase"
	"artics-api/src/middleware"

	"golang.org/x/xerrors"
)

type userService struct {
	auth *firebase.Auth
}

func (us *userService) CreateAuth(ctx context.Context, u *user.User) (string, error) {
	uid, err := us.auth.CreateUser(ctx, u.ID, u.Email, u.Password)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func (us *userService) Auth(ctx context.Context) (string, error) {
	t, err := getToken(ctx)
	if err != nil {
		return "", err
	}

	uid, err := us.auth.VerifyIDToken(ctx, t)
	if err != nil {
		return "", err
	}

	return uid, nil
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
