package content

import (
	"context"
)

type FavoriteRepository interface {
	// Command
	Create(ctx context.Context, f *Favorite) error
	Delete(ctx context.Context, id FavoriteID) error

	// Query
	FindByUserAndContentIDOrNone(ctx context.Context, userID FavoriteUserID, contentID FavoriteContentID) (*QueryFavorite, error)
}
