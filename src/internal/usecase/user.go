package usecase

import (
	"context"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/internal/usecase/validation"

	"golang.org/x/xerrors"
)

// UserUsecase - user usecase
type UserUsecase interface {
	Create(ctx context.Context, r request.CreateUser) (*user.User, error)
	Show(ctx context.Context, id int) (*user.User, error)
	Update(ctx context.Context, r request.UpdateUser) (*user.User, error)
}

type userUsecase struct {
	userRequestValidator validation.UserRequestValidator
	userRepository       user.UserRepository
	userService          user.UserService
}

// NewUserUsecase - generates user usecase
func NewUserUsecase(urv validation.UserRequestValidator, ur user.UserRepository, us user.UserService) UserUsecase {
	return &userUsecase{urv, ur, us}
}

func (uu *userUsecase) Create(ctx context.Context, req request.CreateUser) (*user.User, error) {
	if ves := uu.userRequestValidator.CreateUser(req); len(ves) > 0 {
		err := xerrors.New("Failed to RequestValidation")
		return nil, domain.InvalidRequestValidation.New(err, ves...)
	}

	u := &user.User{
		Nickname:          req.Nickname,
		Email:             req.Email,
		ENcryptedPassword: req.Password,
	}

	uid, err := u.userService.CreateAuth(ctx, u)
	if err != nil {
		return nil, err
	}
	u.UID = uid // Firebase user id
	if err := u.userRepository.Create(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) Show(ctx context.Context, id int) (*user.User, error) {
	if _, err := uu.userService.Auth(ctx); err != nil {
		return nil, domain.Unauthorized.New(err)
	}

	u, err := uu.userRepository.Show(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) Update(ctx context.Context, req request.UpdateUser) (*user.User, error) {
	u, err := uu.userService.Auth(ctx)
	if err != nil {
		return nil, domain.Unauthorized.New(err)
	}

	if ves := uu.userRequestValidator.UpdateUser(req); len(ves) > 0 {
		err := xerrors.New("Failed to RequestValidation")
		return nil, domain.InvalidRequestValidation.New(err, ves...)
	}

	u.Nickname = req.Nickname
	u.Email = req.Email

	if err := uu.userRepository.Update(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}
