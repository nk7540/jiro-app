package command

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
	"context"
	"regexp"

	"github.com/go-playground/validator/v10"
	"golang.org/x/xerrors"
)

type CreateUserHandler struct {
	v  RequestValidator
	ur user.UserRepository
}

func NewCreateUserHandler(v RequestValidator, ur user.UserRepository) CreateUserHandler {
	return CreateUserHandler{v, ur}
}

func (h CreateUserHandler) Handle(ctx pkg.Context, cmd user.CommandCreateUser) error {
	if err := h.v.Validate.RegisterValidation("password", passwordCheck); err != nil {
		return xerrors.Errorf("failed to register validation: %w", err)
	}
	ves := h.v.Run(ctx, cmd)
	vesEmail, err := h.uniqueEmailCheck(ctx, cmd.Email)
	if err != nil {
		return err
	}
	ves = append(ves, vesEmail...)
	if len(ves) > 0 {
		return domain.InvalidDomainValidation.New(pkg.NewValidationError(), ves...)
	}

	uid, err := h.ur.CreateAuth(ctx, cmd)
	if err != nil {
		return domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	}

	u := &user.User{
		UID:   uid,
		Email: user.Email(cmd.Email),
	}

	if err := h.ur.Create(ctx, u); err != nil {
		err = xerrors.Errorf("failed to repository: %w", err)
		return domain.ErrorInDatastore.New(err)
	}

	return nil
}

const (
	passwordString = "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$"
)

var (
	passwordRegex = regexp.MustCompile(passwordString)
)

func passwordCheck(fl validator.FieldLevel) bool {
	return passwordRegex.MatchString(fl.Field().String())
}

func (h CreateUserHandler) uniqueEmailCheck(ctx context.Context, email string) ([]*domain.ValidationError, error) {
	ves := make([]*domain.ValidationError, 0)
	emailUser, err := h.ur.GetByEmailOrNone(ctx, email)
	if err != nil {
		return nil, domain.ErrorInDatastore.New(pkg.NewRepositoryError(err))
	} else if emailUser != nil {
		ve := &domain.ValidationError{
			Field:   "email",
			Message: domain.CustomUniqueMessage,
		}

		ves = append(ves, ve)
	}

	return ves, nil
}
