package content

import (
	"context"
)

type ContentRepository interface {
	Get(ctx context.Context, id int) (*Content, error)
	GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*QueryContent, error)
}
