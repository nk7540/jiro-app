package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"

	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
)

type CreateUser struct {
	Email                string `validate:"required,email,max=256"`
	Password             string `validate:"required,password"`
	PasswordConfirmation string `validate:"required,equalTo=password"`
}

type CreateUserHandler struct {
	userRepository user.UserRepository
}

func NewCreateUserHandler(ur user.UserRepository) CreateUserHandler {
	return CreateUserHandler{ur}
}

func (h CreateUserHandler) Handle(ctx pkg.Context, cmd CreateUser) error {
	email, password, ves := newUser(cmd)
	if len(ves) > 0 {
		err := xerrors.New("failed to domain validation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}
	return h.userRepository.CreateWithPassword(ctx, email, password)
}

func newUser(cmd CreateUser) (user.Email, user.Password, []*domain.ValidationError) {
	validate := validator.New()
	err := validate.Struct(cmd)
	if err == nil {
		return user.Email(cmd.Email), user.Password(cmd.Password), make([]*domain.ValidationError, 0)
	}
}
