package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type UpdateUserHandler struct {
	ur user.UserRepository
}

func NewUpdateUserHandler(ur user.UserRepository) UpdateUserHandler {
	return UpdateUserHandler{ur}
}

func (h UpdateUserHandler) Handle(ctx pkg.Context, cmd user.CommandUpdateUser) (*user.User, error) {
	// Update thumbnail
	thumbnailURL, err := h.ur.UpdateThumbnail(ctx, cmd.Thumbnail)
	if err != nil {
		return nil, domain.ErrorInStorage.New(pkg.NewRepositoryError(err))
	}

	// Update user
	u, err := ctx.CurrentUser()
	if err != nil {
		return nil, err
	}

	u.Nickname = cmd.Nickname
	u.ThumbnailURL = thumbnailURL

	v := NewRequestValidator()
	if ves := v.Run(ctx, u); len(ves) > 0 {
		return nil, domain.InvalidDomainValidation.New(pkg.NewValidationError(), ves...)
	}

	if err := h.ur.Update(ctx, u); err != nil {
		return nil, domain.ErrorInStorage.New(pkg.NewRepositoryError(err))
	}

	return u, nil
}
