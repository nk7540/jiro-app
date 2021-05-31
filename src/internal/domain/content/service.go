package content

import "context"

type ContentService interface {
	GetFavoriteContents(ctx context.Context, userId string, limit int) ([]*Content, error)
}
