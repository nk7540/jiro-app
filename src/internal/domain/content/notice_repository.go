package content

import "context"

type NoticeRepository interface {
	Create(ctx context.Context, n *Notice) error
}
