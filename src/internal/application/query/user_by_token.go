package query

import (
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type UserByTokenHandler struct {
	ur user.UserRepository
}

func NewUserByTokenHandler(ur user.UserRepository) UserByTokenHandler {
	return UserByTokenHandler{ur}
}

func (h UserByTokenHandler) Handle(ctx pkg.Context, tkn string) (*user.User, error) {
	return h.ur.GetByToken(ctx, tkn)
}
