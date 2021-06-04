package service

import (
	"context"
	"io"
	"strings"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/domain/file"
	"artics-api/src/internal/domain/follow"
	"artics-api/src/internal/domain/user"
	"artics-api/src/middleware"

	"golang.org/x/xerrors"
)

type userService struct {
	userDomainValidator user.UserDomainValidator
	userRepository      user.UserRepository
	followRepository    follow.FollowRepository
	contentRepository   content.ContentRepository
	fileRepository      file.FileRepository
}

func NewUserService(
	udv user.UserDomainValidator,
	ur user.UserRepository,
	flwr follow.FollowRepository,
	cr content.ContentRepository,
	flr file.FileRepository,
) user.UserService {
	return &userService{udv, ur, flwr, cr, flr}
}

func (s *userService) Create(ctx context.Context, u *user.User) error {
	ves, err := s.userDomainValidator.Validate(ctx, u)
	if err != nil {
		return err
	}
	vesPassword := s.userDomainValidator.ValidatePassword(ctx, u.Password, u.PasswordConfirmation)
	ves = append(ves, vesPassword...)
	if len(ves) > 0 {
		err := xerrors.New("Failed to DomainValidation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}

	if err := s.userRepository.Create(ctx, u); err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

func (s *userService) Auth(ctx context.Context, tkn string) (*user.User, error) {
	return s.userRepository.GetByToken(ctx, tkn)
}

func (s *userService) Show(ctx context.Context, id int) (*user.User, error) {
	u, err := s.userRepository.Get(ctx, id)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return nil, domain.NotFound.New(err)
	}

	return u, nil
}

func (s *userService) Followings(ctx context.Context, id int) ([]*user.User, error) {
	us, err := s.userRepository.Followings(ctx, id)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return us, nil
}

func (s *userService) Followers(ctx context.Context, id int) ([]*user.User, error) {
	us, err := s.userRepository.Followers(ctx, id)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return nil, domain.ErrorInDatastore.New(err)
	}

	return us, nil
}

func (s *userService) UpdateThumbnail(ctx context.Context, body io.Reader) (string, error) {
	f := &file.File{
		Body: body,
	}
	f, err := s.fileRepository.Save(ctx, f)
	if err != nil {
		err = xerrors.Errorf("Failed to Repository: %w", err)
		return "", domain.ErrorInStorage.New(err)
	}

	return f.Path, nil
}

func (s *userService) Suspend(ctx context.Context, u *user.User) error {
	return s.userRepository.Suspend(ctx, u)
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
