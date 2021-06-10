package content

import "context"

type BrowseRepository interface {
	Save(ctx context.Context, b *Browse) error
}
