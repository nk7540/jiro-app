package repository

import (
	"context"
	"io"

	"artics-api/src/config"
	"artics-api/src/internal/domain/file"
)

type fileRepository struct {
	uploader *config.UploaderConfig
}

func NewFileRepository(uploader *config.UploaderConfig) file.FileRepository {
	return &fileRepository{uploader}
}

func (r *fileRepository) Save(ctx context.Context, body io.Reader) (string, error) {
	output, err := r.uploader.Upload(body)
	if err != nil {
		return "", err
	}
	return output.Location, nil
}
