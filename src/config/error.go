package config

import (
	"artics-api/src/internal/domain"
	"artics-api/src/internal/interface/handler/response"
	"unicode"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type outputLevel int

const (
	debugLevel outputLevel = iota
	infoLevel
	warnLevel
	errorLevel
)

func ErrorHandling(c *fiber.Ctx, err error) error {
	res := getErrorResponse(c, err)

	return c.Status(res.StatusCode).JSON(res)
}

func getErrorResponse(c *fiber.Ctx, err error) *response.ErrorResponse {
	res := response.ErrorResponse{}
	level := infoLevel
	message := ""

	switch getErrorCode(err) {
	// 400
	case domain.InvalidDomainValidation:
		res = *response.BadRequest
		res.ValidationErrors = getValidationErrorsInErrorReponse(c, err)
		message = "Invalid domain validation"
	case domain.InvalidRequestValidation:
		res = *response.BadRequest
		res.ValidationErrors = getValidationErrorsInErrorReponse(c, err)
		message = "Invalid request validation"
	case domain.UnableParseJSON:
		res = *response.BadRequest
		message = "Unable parse json"
	case domain.NotEqualRequestWithDatastore:
		res = *response.BadRequest
		message = "Invalid request validation"
	// 401
	case domain.Unauthorized:
		res = *response.Unauthorized
		message = "Unauthorized"
	// 403
	case domain.Forbidden:
		res = *response.Forbidden
		message = "Forbidden"
	// 404
	case domain.NotFound:
		res = *response.NotFound
		message = "Not found"
	// 409
	case domain.AlreadyExistsInDatastore:
		res = *response.AlreadyExists
		res.ValidationErrors = getValidationErrorsInErrorReponse(c, err)
		message = "Already exists request"
	// 500
	case domain.ErrorInDatastore:
		res = *response.InternalServerError
		level = warnLevel
		message = "Error in datastore"
	case domain.ErrorInStorage:
		res = *response.InternalServerError
		level = warnLevel
		message = "Error in storage"
	default:
		res = *response.InternalServerError
		level = errorLevel
		message = "Internal server error"
	}
	p := c.Locals("i18n").(I18nConfig)
	res.Message = p.Sprintf(res.Message)

	res.ErrorCode = getErrorCode(err)
	logging(level, message, err, &res)

	return &res
}

func logging(level outputLevel, message string, err error, res *response.ErrorResponse) {
	fields := log.Fields{
		"message":   message,
		"errorCode": res.ErrorCode,
	}

	if len(res.ValidationErrors) > 0 {
		fields["validationErrors"] = res.ValidationErrors
	}

	switch level {
	case debugLevel:
		log.WithFields(fields).Debug(err.Error())
	case infoLevel:
		log.WithFields(fields).Info(err.Error())
	case warnLevel:
		log.WithFields(fields).Info(err.Error())
	default:
		log.WithFields(fields).Error(err.Error())
	}
}

func getErrorCode(err error) domain.ErrorCode {
	if e, ok := err.(domain.CustomError); ok {
		return e.Code()
	}

	return domain.Unknown
}

func getValidationErrorsInErrorReponse(c *fiber.Ctx, err error) response.ValidationErrors {
	p := c.Locals("i18n").(I18nConfig)
	if e, ok := err.(domain.CustomError); ok {
		ves := response.ValidationErrors{}
		for _, ve := range e.Validations() {
			a := []rune(ve.Field)
			a[0] = unicode.ToLower(a[0])
			field := string(a)
			ves[field] = p.Sprintf(field) + p.Sprintf(ve.Message)
		}

		return ves
	}

	return response.ValidationErrors{}
}
