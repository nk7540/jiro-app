package browse

import "context"

type BrowseService interface {
	Save(ctx context.Context, b *Browse) error
}
