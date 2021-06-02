package content

import "context"

type ContentService interface {
	GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*Content, error)
}
