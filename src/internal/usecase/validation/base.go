package validation

import (
	"context"
	"reflect"

	"artics-api/src/config"
	"artics-api/src/internal/domain"
	"artics-api/src/pkg"

	"github.com/go-playground/validator/v10"
)

// RequestValidator - base validator interface
type RequestValidator interface {
	Run(ctx context.Context, i interface{}) []*domain.ValidationError
}

type requestValidator struct {
	validate validator.Validate
}

// NewRequestValidator - generate base validator
func NewRequestValidator() RequestValidator {
	validate := validator.New()

	return &requestValidator{
		validate: *validate,
	}
}

// Run - runs a validation
func (rv *requestValidator) Run(ctx context.Context, i interface{}) []*domain.ValidationError {
	err := rv.validate.Struct(i)
	if err == nil {
		return make([]*domain.ValidationError, 0)
	}

	errors := err.(validator.ValidationErrors)
	validationErrors := make([]*domain.ValidationError, len(errors))

	rt := reflect.ValueOf(i).Elem().Type()

	for i, v := range errors {
		errorField, _ := rt.FieldByName(v.Field())
		errorFieldName := errorField.Tag.Get("json")
		errorMessage := ""

		c := ctx.(pkg.Context)
		p := c.Locals("i18n").(config.I18nConfig)
		switch v.Tag() {
		case domain.EqFieldTag:
			eqField, _ := rt.FieldByName(v.Param())
			errorMessage = validationMessage(&p, v.Tag(), eqField.Tag.Get("label"))
		default:
			errorMessage = validationMessage(&p, v.Tag(), v.Param())
		}

		validationErrors[i] = &domain.ValidationError{
			Field:   errorFieldName,
			Message: errorMessage,
		}
	}

	return validationErrors
}

func validationMessage(p *config.I18nConfig, tag string, options ...string) string {
	switch tag {
	case domain.RequiredTag:
		return p.Sprintf(domain.RequiredMessage)
	case domain.EqFieldTag:
		return p.Sprintf(domain.EqFieldMessage, options[0])
	case domain.MinTag:
		return p.Sprintf(domain.MinMessage, options[0])
	case domain.MaxTag:
		return p.Sprintf(domain.MaxMessage, options[0])
	case domain.EmailTag:
		return p.Sprintf(domain.EmailMessage)
	case domain.PasswordTag:
		return p.Sprintf(domain.PasswordMessage)
	case domain.UniqueTag:
		return p.Sprintf(domain.UniqueMessage)
	default:
		return ""
	}
}
