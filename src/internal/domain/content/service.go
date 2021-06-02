package content

import "context"

type ContentService interface {
	Get(ctx context.Context, id int) (*Content, error)
	GetFavoriteContents(ctx context.Context, userId int, limit int) ([]*Content, error)
}
