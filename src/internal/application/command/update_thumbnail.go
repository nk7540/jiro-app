package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
	"context"
)

type UpdateThumbnailHandler struct {
	ur user.UserRepository
}

func NewUpdateThumbnailHandler(ur user.UserRepository) UpdateThumbnailHandler {
	return UpdateThumbnailHandler{ur}
}

func (h UpdateThumbnailHandler) Handle(ctx context.Context, thumbnail user.Thumbnail) (user.ThumbnailURL, error) {
	thumbnailURL, err := h.ur.UpdateThumbnail(ctx, thumbnail)
	if err != nil {
		return "", domain.ErrorInStorage.New(pkg.NewRepositoryError(err))
	}

	return thumbnailURL, nil
}
