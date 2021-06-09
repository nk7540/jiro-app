package command

import (
	"artics-api/src/internal/domain/file"
	"context"
	"io"
)

type UpdateThumbnailHandler struct {
	fileRepository file.FileRepository
}

func NewUpdateThumbnailHandler(fr file.FileRepository) UpdateThumbnailHandler {
	return UpdateThumbnailHandler{fr}
}

func (h UpdateThumbnailHandler) Handle(ctx context.Context, body io.Reader) (string, error) {
	return h.fileRepository.Save(ctx, body)
}
