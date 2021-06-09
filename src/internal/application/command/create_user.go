package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"

	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
)

type CreateUserHandler struct {
	userRepository user.UserRepository
}

func NewCreateUserHandler(ur user.UserRepository) CreateUserHandler {
	return CreateUserHandler{ur}
}

func (h CreateUserHandler) Handle(ctx pkg.Context, cmd user.CommandCreateUser) error {
	if ves := validate(cmd); len(ves) > 0 {
		err := xerrors.New("failed to domain validation")
		return domain.InvalidDomainValidation.New(err, ves...)
	}
	uid, err := h.userRepository.CreateAuth(ctx, cmd)
	if err != nil {
		return err
	}
	u := &user.User{
		UID:   uid,
		Email: user.Email(cmd.Email),
	}
	return h.userRepository.Create(ctx, u)
}

func validate(cmd user.CommandCreateUser) []*domain.ValidationError {
	validate := validator.New()
	err := validate.Struct(cmd)
	if err == nil {
		return make([]*domain.ValidationError, 0)
	}
}
