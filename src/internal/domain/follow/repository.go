package follow

import (
	"artics-api/src/internal/domain/user"
	"context"
)

type FollowRepository interface {
	// Command
	Create(ctx context.Context, follow *Follow) error
	Delete(ctx context.Context, id ID) error

	// Query
	GetByUserIDs(ctx context.Context, followingID FollowingID, followerID FollowerID) (*QueryFollow, error)
	FollowingCount(ctx context.Context, userID user.ID) (int, error)
}
