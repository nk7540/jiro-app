package command

import (
	"artics-api/src/internal/domain/user"
	"context"
	"fmt"
)

type SuspendUserHandler struct {
	userRepository user.UserRepository
}

func NewSuspendUserHandler(ur user.UserRepository) SuspendUserHandler {
	return SuspendUserHandler{ur}
}

func (h SuspendUserHandler) Handle(ctx context.Context, u *user.User) error {
	// prevEmail := u.Email

	u.Status = user.Suspended
	u.Nickname = ""
	u.Email = user.Email(fmt.Sprintf("leaved+user%s@artics.jp", u.ID))
	u.ThumbnailURL = ""

	// @TODO validation
	if err := h.userRepository.Update(ctx, u); err != nil {
		return err
	}

	// @TODO notify suspended
	return h.userRepository.DeleteAuth(ctx, u.UID)
}
