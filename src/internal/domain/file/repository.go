package file

import "context"

type FileRepository interface {
	Save(ctx context.Context, f *File) (*File, error)
}
