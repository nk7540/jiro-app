package content

import (
	"context"
)

type ContentRepository interface {
	GetByIDs(ctx context.Context, ids []string) ([]*Content, error)
}
