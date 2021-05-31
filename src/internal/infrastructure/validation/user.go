package validation

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/domain/user"
	"artics-api/src/lib/i18n"
	"artics-api/src/middleware"
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

func (udv *userDomainValidator) Validate(ctx context.Context, u *user.User) []*domain.ValidationError {
	c, _ := middleware.GinContextFromContext(ctx)
	p := i18n.NewI18nPrinter(c.GetHeader("Accept-Language"))
	ves := make([]*domain.ValidationError, 0)

	emailUser, _ := udv.ur.GetByEmail(ctx, u.Email)
	if emailUser != nil && emailUser.ID != u.ID {
		ve := &domain.ValidationError{
			Field:   "email",
			Message: p.Sprintf(domain.CustomUniqueMessage),
		}

		ves = append(ves, ve)
	}

	return ves
}

func (udv *userDomainValidator) ValidatePassword(ctx context.Context, password string, passwordConfirmation string) []*domain.ValidationError {
	c, _ := middleware.GinContextFromContext(ctx)
	p := i18n.NewI18nPrinter(c.GetHeader("Accept-Language"))
	ves := make([]*domain.ValidationError, 0)

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
