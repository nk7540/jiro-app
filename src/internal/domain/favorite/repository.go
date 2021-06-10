package favorite

import (
	"context"
)

type FavoriteRepository interface {
	// Command
	Create(ctx context.Context, f *Favorite) error
	Delete(ctx context.Context, id ID) error

	// Query
	FindByUserAndContentID(ctx context.Context, userID UserID, contentID ContentID) (*QueryFavorite, error)
}
