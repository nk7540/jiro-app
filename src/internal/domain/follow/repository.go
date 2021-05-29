package follow

import (
	"context"
)

type FollowRepository interface {
	FollowingCount(ctx context.Context, id int) (int, error)
	FollowerCount(ctx context.Context, id int) (int, error)
	Create(ctx context.Context, follow *Follow) error
	Delete(ctx context.Context, follow *Follow) error
}

