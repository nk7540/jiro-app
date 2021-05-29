package validation

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/usecase/request"
)

// UserRequestValidator - user validator interface
type UserRequestValidator interface {
	CreateUser(req *request.CreateUser) []*domain.ValidationError
	UpdateUser(req *request.UpdateUser) []*domain.ValidationError
}

type userRequestValidator struct {
	validator RequestValidator
}

// NewUserRequestValidator - generates user validator
func NewUserRequestValidator() UserRequestValidator {
	rv := NewRequestValidator()

	return &userRequestValidator{
		validator: rv,
	}
}

func (urv *userRequestValidator) CreateUser(req *request.CreateUser) []*domain.ValidationError {
	return urv.validator.Run(req)
}

func (urv *userRequestValidator) UpdateUser(req *request.UpdateUser) []*domain.ValidationError {
	return urv.validator.Run(req)
}
