package query

import (
	"artics-api/src/internal/domain/user"
	"context"
)

type FollowersHandler struct {
	userRepository user.UserRepository
}

func NewFollowersHandler(ur user.UserRepository) FollowersHandler {
	return FollowersHandler{ur}
}

func (h FollowersHandler) Handle(ctx context.Context, userID user.UserID) (*user.QueryUsers, error) {
	return h.userRepository.Followers(ctx, userID)
}
