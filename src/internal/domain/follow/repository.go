package follow

import (
	"context"
)

type FollowRepository interface {
	FollowingCount(ctx context.Context, id string) (int, error)
	FollowerCount(ctx context.Context, id string) (int, error)
	Create(ctx context.Context, follow *Follow) error
	Delete(ctx context.Context, follow *Follow) error
}

