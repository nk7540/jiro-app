package content

import (
	"context"
)

type ContentRepository interface {
	GetFavoriteContents(ctx context.Context, userId string, limit int) ([]*Content, error)
}
