package request

// CreateUser - request from CreateUser
type CreateUser struct {
	Nickname             string `json:"nickname" validate:"required,max=256"`
	Email                string `json:"email" validate:"required,email,max=256"`
	Password             string `json:"password" validate:"password,required,min=6,max=32"`
	PasswordConfirmation string `json:"passwordConfirmation" validate:"required,eqfield=Password"`
}

// UpdateUser - request from UpdateUser
type UpdateUser struct {
	Nickname string `json:"nickname" validate:"max=256"`
	Email    string `json:"email" validate:"email,max=256"`
}
