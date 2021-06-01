package follow

import (
	"context"
)

type FollowRepository interface {
	Create(ctx context.Context, follow *Follow) error
	Delete(ctx context.Context, follow *Follow) error
}
