package favorite

import "context"

type FavoriteService interface {
	Create(ctx context.Context, f *Favorite) error
	Delete(ctx context.Context, f *Favorite) error
}
