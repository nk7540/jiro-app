package favorite

import (
	"context"
)

type FavoriteRepository interface {
	FavoriteContentIds(ctx context.Context, userId int) (int[], error)
}
