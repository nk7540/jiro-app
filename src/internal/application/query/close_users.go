package query

import (
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
)

type CloseUsersHandler struct {
	ur user.UserRepository
}

func NewCloseUsersHandler(ur user.UserRepository) CloseUsersHandler {
	return CloseUsersHandler{ur}
}

func (h CloseUsersHandler) Handle(ctx pkg.Context, userID user.UserID) ([]*user.QueryUsers, error) {

}
