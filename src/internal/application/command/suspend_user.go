package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
	"context"

	"golang.org/x/xerrors"
)

type SuspendUserHandler struct {
	ur user.UserRepository
}

func NewSuspendUserHandler(ur user.UserRepository) SuspendUserHandler {
	return SuspendUserHandler{ur}
}

func (h SuspendUserHandler) Handle(ctx context.Context, u *user.User) error {
	// prevEmail := u.Email
	u.Suspend()

	v := NewRequestValidator()
	if ves := v.Run(ctx, u); len(ves) > 0 {
		return xerrors.New("failed to suspended user validation")
	}

	if err := h.ur.Update(ctx, u); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	// @TODO notify suspended
	if err := h.ur.DeleteAuth(ctx, u.UID); err != nil {
		return domain.ErrorInStorage.New(err)
	}

	return nil
}
