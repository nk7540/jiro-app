package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type UpdateUserHandler struct {
	v  RequestValidator
	ur user.UserRepository
}

func NewUpdateUserHandler(v RequestValidator, ur user.UserRepository) UpdateUserHandler {
	return UpdateUserHandler{v, ur}
}

func (h UpdateUserHandler) Handle(ctx pkg.Context, cmd user.CommandUpdateUser) error {
	u := cmd.User
	u.Nickname = user.Nickname(cmd.Nickname)
	u.ThumbnailURL = user.ThumbnailURL(cmd.ThumbnailURL)

	if ves := h.v.Run(ctx, u); len(ves) > 0 {
		return domain.InvalidDomainValidation.New(pkg.NewValidationError(), ves...)
	}

	return h.ur.Update(ctx, u)
}
