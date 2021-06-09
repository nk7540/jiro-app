package query

import (
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type GetUserHandler struct {
	userRepository user.UserRepository
}

func NewGetUserHandler(ur user.UserRepository) GetUserHandler {
	return GetUserHandler{ur}
}

func (h GetUserHandler) Handle(ctx pkg.Context, id int) (*user.QueryDetailUser, error) {
	return h.userRepository.Get(ctx, id)
}
