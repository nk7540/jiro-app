package repository

import (
	"context"

	"artics-api/src/internal/domain/file"
	"artics-api/src/lib/awssdk"
)

type fileRepository struct {
	au *awssdk.Uploader
}

func NewFileRepository(au *awssdk.Uploader) file.FileRepository {
	return &fileRepository{au}
}

func (r *fileRepository) Save(ctx context.Context, f *file.File) (*file.File, error) {
	output, err := r.au.Upload(f.Body)
	if err != nil {
		return nil, err
	}
	f.Path = output.Location

	return f, nil
}
