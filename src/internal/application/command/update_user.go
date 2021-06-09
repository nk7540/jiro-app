package command

import (
	"artics-api/src/internal/domain/user"
	"context"
)

type UpdateUserHandler struct {
	userRepository user.UserRepository
}

func NewUpdateUserHandler(ur user.UserRepository) UpdateUserHandler {
	return UpdateUserHandler{ur}
}

func (h UpdateUserHandler) Handle(ctx context.Context, cmd user.CommandUpdateUser) error {
	// @TODO validation
	return h.userRepository.Update(ctx, cmd)
}
