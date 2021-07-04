package content

import "context"

type CommentRepository interface {
	Create(ctx context.Context, c *Comment) (CommentID, error)
}
