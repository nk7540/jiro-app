package user

import (
	"context"
)

type FollowRepository interface {
	// Command
	Create(ctx context.Context, follow *Follow) error
	Delete(ctx context.Context, id FollowID) error

	// Query
	GetByUserIDs(ctx context.Context, followingID FollowingID, followerID FollowerID) (*QueryFollow, error)
	FollowingCount(ctx context.Context, userID UserID) (int, error)
}
