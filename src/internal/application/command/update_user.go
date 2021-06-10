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

func (h UpdateUserHandler) Handle(ctx pkg.Context, cmd user.CommandUpdateUser) error {
	u, err := ctx.CurrentUser()
	if err != nil {
		return err
	}

	u.Nickname = user.Nickname(cmd.Nickname)
	u.ThumbnailURL = user.ThumbnailURL(cmd.ThumbnailURL)

	v := NewRequestValidator()
	if ves := v.Run(ctx, u); len(ves) > 0 {
		return domain.InvalidDomainValidation.New(pkg.NewValidationError(), ves...)
	}

	return h.ur.Update(ctx, u)
}
