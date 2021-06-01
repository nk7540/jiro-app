package request

import "mime/multipart"

// CreateUser - request from CreateUser
type CreateUser struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

// UpdateUser - request from UpdateUser
type UpdateUser struct {
	Nickname  string                `form:"nickname"`
	Email     string                `form:"email"`
	Thumbnail *multipart.FileHeader `form:"thumbnail"`
}
