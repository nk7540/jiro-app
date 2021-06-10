package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/favorite"
	"artics-api/src/pkg"

	"golang.org/x/xerrors"
)

type LikeHandler struct {
	fr favorite.FavoriteRepository
}

func NewLikeHandler(fr favorite.FavoriteRepository) LikeHandler {
	return LikeHandler{fr}
}

func (h LikeHandler) Handle(ctx pkg.Context, cmd favorite.CommandLike) error {
	if qf, err := h.fr.FindByUserAndContentID(ctx, cmd.UserID, cmd.ContentID); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if qf != nil {
		return xerrors.New("already in favorite")
	}

	f := &favorite.Favorite{
		UserID:    cmd.UserID,
		ContentID: cmd.ContentID,
	}

	if err := h.fr.Create(ctx, f); err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	return nil
}
