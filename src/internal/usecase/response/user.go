package response

import (
	"context"

	"artics-api/src/internal/domain/user"
)

type CreateUser struct {
	ResultCode string `json:"resultCode"`
}

type ShowUser struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}
