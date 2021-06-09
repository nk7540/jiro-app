package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
	"context"
	"fmt"
)

type SuspendUserHandler struct {
	v  RequestValidator
	ur user.UserRepository
}

func NewSuspendUserHandler(v RequestValidator, ur user.UserRepository) SuspendUserHandler {
	return SuspendUserHandler{v, ur}
}

func (h SuspendUserHandler) Handle(ctx context.Context, u *user.User) error {
	// prevEmail := u.Email

	u.Status = user.Suspended
	u.Nickname = ""
	u.Email = user.Email(fmt.Sprintf("leaved+user%s@artics.jp", u.ID))
	u.ThumbnailURL = ""

	if ves := h.v.Run(ctx, u); len(ves) > 0 {
		return domain.InvalidDomainValidation.New(pkg.NewValidationError(), ves...)
	}

	if err := h.ur.Update(ctx, u); err != nil {
		return err
	}

	// @TODO notify suspended
	return h.ur.DeleteAuth(ctx, u.UID)
}
