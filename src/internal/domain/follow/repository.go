package follow

import (
	"artics-api/src/internal/domain/user"
	"context"
)

type FollowRepository interface {
	Create(ctx context.Context, follow *Follow) error
	FollowingCount(ctx context.Context, userID user.ID) (int, error)
	Delete(ctx context.Context, follow *Follow) error
}
