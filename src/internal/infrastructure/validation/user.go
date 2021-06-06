package validation

import (
	"artics-api/src/config"
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/pkg"
	"context"
	"regexp"
)

type userDomainValidator struct {
	ur user.UserRepository
}

func NewUserDomainValidator(ur user.UserRepository) user.UserDomainValidator {
	return &userDomainValidator{ur}
}

const (
	passwordString = "^[a-zA-Z0-9_!@#$_%^&*.?()-=+]*$"
)

var (
	passwordRegexp = regexp.MustCompile(passwordString)
)

func (udv *userDomainValidator) Validate(ctx context.Context, u *user.User) ([]*domain.ValidationError, error) {
	c := ctx.(pkg.Context)
	p := c.Locals("i18n").(config.I18nConfig)
	ves := make([]*domain.ValidationError, 0)

	emailUser, err := udv.ur.GetByEmailOrNone(ctx, u.Email)
	if err != nil {
		return nil, domain.ErrorInDatastore.New(err)
	}
	if emailUser != nil && emailUser.ID != u.ID {
		ve := &domain.ValidationError{
			Field:   "email",
			Message: p.Sprintf(domain.CustomUniqueMessage),
		}

		ves = append(ves, ve)
	}

	return ves, nil
}

func (udv *userDomainValidator) ValidatePassword(ctx context.Context, password string, passwordConfirmation string) []*domain.ValidationError {
	c := ctx.(pkg.Context)
	p := c.Locals("i18n").(config.I18nConfig)
	ves := make([]*domain.ValidationError, 0)

	if password == "" {
		ve := &domain.ValidationError{
			Field:   "password",
			Message: p.Sprintf(domain.RequiredMessage),
		}

		ves = append(ves, ve)
	}

	passwordLength := len(password)
	if passwordLength < 6 {
		ve := &domain.ValidationError{
			Field:   "password",
			Message: p.Sprintf(domain.MinMessage, "6"),
		}

		ves = append(ves, ve)
	} else if 32 <= passwordLength {
		ve := &domain.ValidationError{
			Field:   "password",
			Message: p.Sprintf(domain.MaxMessage, "32"),
		}

		ves = append(ves, ve)
	}

	formatValid := passwordRegexp.MatchString(password)
	if !formatValid {
		ve := &domain.ValidationError{
			Field:   "password",
			Message: p.Sprintf(domain.PasswordMessage),
		}

		ves = append(ves, ve)
	}

	if password != passwordConfirmation {
		ve := &domain.ValidationError{
			Field:   "passwordConfirmation",
			Message: p.Sprintf(domain.PasswordConfirmationMessage),
		}

		ves = append(ves, ve)
	}

	return ves
}
