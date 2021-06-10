package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/pkg"

	"golang.org/x/xerrors"
)

type UnlikeHandler struct {
	fr content.FavoriteRepository
}

func NewUnlikeHandler(fr content.FavoriteRepository) UnlikeHandler {
	return UnlikeHandler{fr}
}

func (h UnlikeHandler) Handle(ctx pkg.Context, cmd content.CommandUnlike) error {
	qf, err := h.fr.FindByUserAndContentIDOrNone(ctx, cmd.UserID, cmd.ContentID)
	if err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if qf == nil {
		return xerrors.New("not in favorite")
	}

	if err := h.fr.Delete(ctx, content.FavoriteID(qf.ID)); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
