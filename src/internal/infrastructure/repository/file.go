package repository

import (
	"context"

	"artics-api/src/config"
	"artics-api/src/internal/domain/file"
)

type fileRepository struct {
	uploader *config.UploaderConfig
}

func NewFileRepository(uploader *config.UploaderConfig) file.FileRepository {
	return &fileRepository{uploader}
}

func (r *fileRepository) Save(ctx context.Context, f *file.File) (*file.File, error) {
	output, err := r.uploader.Upload(f.Body)
	if err != nil {
		return nil, err
	}
	f.Path = output.Location

	return f, nil
}
