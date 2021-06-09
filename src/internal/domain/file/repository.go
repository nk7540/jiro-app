package file

import (
	"context"
	"io"
)

type FileRepository interface {
	Save(ctx context.Context, body io.Reader) (string, error)
}
