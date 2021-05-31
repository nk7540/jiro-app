package usecase

import (
	"context"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/follow"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/internal/usecase/validation"

	"golang.org/x/xerrors"
)

// UserUsecase - user usecase
type UserUsecase interface {
	Create(ctx context.Context, r *request.CreateUser) error
	Show(ctx context.Context, id string) (*user.User, error)
	Update(ctx context.Context, r *request.UpdateUser) (*user.User, error)
	Suspend(ctx context.Context) error
}

type userUsecase struct {
	RequestValidator validation.RequestValidator
	userRepository   user.UserRepository
	userService      user.UserService
	followRepository follow.FollowRepository
}

// NewUserUsecase - generates user usecase
func NewUserUsecase(rv validation.RequestValidator, ur user.UserRepository, us user.UserService, fr follow.FollowRepository) UserUsecase {
	return &userUsecase{rv, ur, us, fr}
}

func (uu *userUsecase) Create(ctx context.Context, req *request.CreateUser) error {
	u := &user.User{
		Status:               "provisional",
		Email:                req.Email,
		Password:             req.Password,
		PasswordConfirmation: req.PasswordConfirmation,
	}

	if ves := uu.RequestValidator.Run(ctx, u); len(ves) > 0 {
		err := xerrors.New("Failed to RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	return uu.userService.Create(ctx, u)
}

func (uu *userUsecase) Show(ctx context.Context, id string) (*user.User, error) {
	if _, err := uu.userService.Auth(ctx); err != nil {
		return nil, domain.Unauthorized.New(err)
	}

	return uu.userService.Show(ctx, id)
}

func (uu *userUsecase) Update(ctx context.Context, req *request.UpdateUser) (*user.User, error) {
	u, err := uu.userService.Auth(ctx)
	if err != nil {
		return nil, domain.Unauthorized.New(err)
	}

	u.Nickname = req.Nickname
	u.Email = req.Email

	if ves := uu.RequestValidator.Run(ctx, u); len(ves) > 0 {
		err := xerrors.New("Failed to RequestValidation")
		return nil, domain.InvalidRequestValidation.New(err, ves...)
	}

	if err := uu.userRepository.Update(ctx, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (uu *userUsecase) Suspend(ctx context.Context) error {
	u, err := uu.userService.Auth(ctx)
	if err != nil {
		return domain.Unauthorized.New(err)
	}

	u.Status = "suspended"

	if ves := uu.RequestValidator.Run(ctx, u); len(ves) > 0 {
		err := xerrors.New("Failed to RequestValidation")
		return domain.InvalidRequestValidation.New(err, ves...)
	}

	return uu.userService.Suspend(ctx, u)
}
