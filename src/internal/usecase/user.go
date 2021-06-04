package usecase

import (
	"context"

	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/internal/domain/user"
	"artics-api/src/internal/usecase/request"
	"artics-api/src/internal/usecase/response"
	"artics-api/src/internal/usecase/validation"

	"golang.org/x/xerrors"
)

// UserUsecase - user usecase
type UserUsecase interface {
	Create(ctx context.Context, r *request.CreateUser) error
	Auth(ctx context.Context, tkn string) (*user.User, error)
	Show(ctx context.Context, id int) (*response.ShowUser, error)
	Followings(ctx context.Context, id int) (*response.Users, error)
	Followers(ctx context.Context, id int) (*response.Users, error)
	Update(ctx context.Context, r *request.UpdateUser) (*response.UpdateUser, error)
	Suspend(ctx context.Context) error
}

type userUsecase struct {
	RequestValidator validation.RequestValidator
	userRepository   user.UserRepository
	userService      user.UserService
	contentService   content.ContentService
}

// NewUserUsecase - generates user usecase
func NewUserUsecase(
	rv validation.RequestValidator,
	ur user.UserRepository,
	us user.UserService,
	cs content.ContentService,
) UserUsecase {
	return &userUsecase{rv, ur, us, cs}
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

func (uu *userUsecase) Auth(ctx context.Context, tkn string) (*user.User, error) {
	return uu.userService.Auth(ctx, tkn)
}

func (uu *userUsecase) Show(ctx context.Context, id int) (*response.ShowUser, error) {
	if _, err := uu.userService.Auth(ctx); err != nil {
		return nil, domain.Unauthorized.New(err)
	}

	u, err := uu.userService.Show(ctx, id)
	if err != nil {
		return nil, err
	}
	favoriteContents, err := uu.contentService.GetFavoriteContents(ctx, id, 3)
	if err != nil {
		return nil, err
	}

	resFavoriteContents := make([]*response.Content, len(favoriteContents))
	for i, c := range favoriteContents {
		resFavoriteContents[i] = &response.Content{
			ID:    c.ID,
			Title: c.Title,
		}
	}

	res := &response.ShowUser{
		ID:               u.ID,
		Nickname:         u.Nickname,
		Email:            u.Email,
		Followingcount:   u.FollowingCount,
		Followercount:    u.FollowerCount,
		FavoriteContents: resFavoriteContents,
	}

	return res, nil
}

func (uu *userUsecase) Followings(ctx context.Context, id int) (*response.Users, error) {
	us, err := uu.userService.Followings(ctx, id)
	if err != nil {
		return nil, err
	}

	resUsers := make([]*response.User, len(us))
	for i, u := range us {
		resUsers[i] = &response.User{
			ID:       u.ID,
			Nickname: u.Nickname,
		}
	}
	res := &response.Users{resUsers}

	return res, nil
}

func (uu *userUsecase) Followers(ctx context.Context, id int) (*response.Users, error) {
	us, err := uu.userService.Followers(ctx, id)
	if err != nil {
		return nil, err
	}

	resUsers := make([]*response.User, len(us))
	for i, u := range us {
		resUsers[i] = &response.User{
			ID:       u.ID,
			Nickname: u.Nickname,
		}
	}
	res := &response.Users{resUsers}

	return res, nil
}

func (uu *userUsecase) Update(ctx context.Context, req *request.UpdateUser) (*response.UpdateUser, error) {
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

	thumbnail, err := req.Thumbnail.Open()
	if err != nil {
		return nil, domain.UnableParseFile.New(err)
	}
	thumbnailURL, err := uu.userService.UpdateThumbnail(ctx, thumbnail)
	if err != nil {
		return nil, err
	}

	u.ThumbnailURL = thumbnailURL

	if err := uu.userRepository.Update(ctx, u); err != nil {
		return nil, err
	}

	res := &response.UpdateUser{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
	}

	return res, nil
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
