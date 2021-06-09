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

func (h FollowersHandler) Handle(ctx context.Context, userID int) ([]*user.QueryUser, error) {
	return h.userRepository.Followers(ctx, userID)
}
