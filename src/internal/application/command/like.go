package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/content"
	"artics-api/src/pkg"

	"golang.org/x/xerrors"
)

type LikeHandler struct {
	fr content.FavoriteRepository
}

func NewLikeHandler(fr content.FavoriteRepository) LikeHandler {
	return LikeHandler{fr}
}

func (h LikeHandler) Handle(ctx pkg.Context, cmd content.CommandLike) error {
	if qf, err := h.fr.FindByUserAndContentIDOrNone(ctx, cmd.UserID, cmd.ContentID); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if qf != nil {
		return xerrors.New("already in favorite")
	}

	f := &content.Favorite{
		UserID:    cmd.UserID,
		ContentID: cmd.ContentID,
	}

	if err := h.fr.Create(ctx, f); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
