package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/follow"
	"artics-api/src/pkg"

	"golang.org/x/xerrors"
)

type UnfollowHandler struct {
	fr follow.FollowRepository
}

func NewUnfollowHandler(fr follow.FollowRepository) UnfollowHandler {
	return UnfollowHandler{fr}
}

func (h UnfollowHandler) Handle(ctx pkg.Context, cmd follow.CommandUnfollow) error {
	qf, err := h.fr.GetByUserIDs(ctx, cmd.FollowingID, cmd.FollowerID)
	if err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if qf == nil {
		return xerrors.New("not following")
	}

	if err := h.fr.Delete(ctx, follow.ID(qf.ID)); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
