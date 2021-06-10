package content

import (
	"context"
)

type ContentRepository interface {
	GetOrNone(ctx context.Context, id ContentID) (*QueryDetailContent, error)
	GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*QueryContent, error)
}
