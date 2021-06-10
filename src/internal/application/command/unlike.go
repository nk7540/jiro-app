package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/favorite"
	"artics-api/src/pkg"

	"golang.org/x/xerrors"
)

type UnlikeHandler struct {
	fr favorite.FavoriteRepository
}

func NewUnlikeHandler(fr favorite.FavoriteRepository) UnlikeHandler {
	return UnlikeHandler{fr}
}

func (h UnlikeHandler) Handle(ctx pkg.Context, cmd favorite.CommandUnlike) error {
	qf, err := h.fr.FindByUserAndContentID(ctx, cmd.UserID, cmd.ContentID)
	if err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if qf == nil {
		return xerrors.New("not in favorite")
	}

	if err := h.fr.Delete(ctx, favorite.ID(qf.ID)); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
