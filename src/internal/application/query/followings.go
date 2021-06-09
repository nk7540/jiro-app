package query

import (
	"artics-api/src/internal/domain/user"
	"context"
)

type FollowingsHandler struct {
	userRepository user.UserRepository
}

func NewFollowingsHandler(ur user.UserRepository) FollowingsHandler {
	return FollowingsHandler{ur}
}

func (h FollowingsHandler) Handle(ctx context.Context, userID int) ([]*user.QueryUser, error) {
	return h.userRepository.Followings(ctx, userID)
}
