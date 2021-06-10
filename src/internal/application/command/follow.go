package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"

	"golang.org/x/xerrors"
)

type FollowHandler struct {
	fr user.FollowRepository
}

func NewFollowHandler(fr user.FollowRepository) FollowHandler {
	return FollowHandler{fr}
}

func (h FollowHandler) Handle(ctx pkg.Context, cmd user.CommandFollow) error {
	if qf, err := h.fr.GetByUserIDs(ctx, cmd.FollowingID, cmd.FollowerID); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if qf != nil {
		return xerrors.New("already following")
	}

	f := &user.Follow{
		FollowingID: cmd.FollowingID,
		FollowerID:  cmd.FollowerID,
	}

	if err := h.fr.Create(ctx, f); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
