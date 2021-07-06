package notice

import "context"

type NoticeRepository interface {
	Create(ctx context.Context, n *Notice) error
	List(ctx context.Context, qry QueryArgListNotice) (*QueryNotices, error)
}
