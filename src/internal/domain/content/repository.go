package content

import (
	"context"
)

type ContentRepository interface {
	GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*Content, error)
}
