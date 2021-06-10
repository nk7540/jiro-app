package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/follow"
	"artics-api/src/pkg"

	"golang.org/x/xerrors"
)

type FollowHandler struct {
	fr follow.FollowRepository
}

func NewFollowHandler(fr follow.FollowRepository) FollowHandler {
	return FollowHandler{fr}
}

func (h FollowHandler) Handle(ctx pkg.Context, cmd follow.CommandFollow) error {
	if qf, err := h.fr.GetByUserIDs(ctx, cmd.FollowingID, cmd.FollowerID); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if qf != nil {
		return xerrors.New("already following")
	}

	f := &follow.Follow{
		FollowingID: cmd.FollowingID,
		FollowerID:  cmd.FollowerID,
	}

	if err := h.fr.Create(ctx, f); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
