package domain

type CustomError struct {
	ErrorCode        ErrorCode
	Value            error
	ValidationErrors []*ValidationError
}

type ValidationError struct {
	Field   string
	Message string
}

type ErrorCode uint

const (
	Unknown ErrorCode = iota
	Unauthorized
	Forbidden
	NotFound
	InvalidDomainValidation
	InvalidRequestValidation
	UnableParseJSON
	UnableParseFormData
	UnableParseFile
	ErrorInDatastore
	AlreadyExistsInDatastore
	NotEqualRequestWithDatastore
	ErrorInStorage
)

func (ec ErrorCode) New(err error, ves ...*ValidationError) error {
	return CustomError{
		ErrorCode:        ec,
		Value:            err,
		ValidationErrors: ves,
	}
}

func (ce CustomError) Error() string {
	return ce.Value.Error()
}

func (ce CustomError) Code() ErrorCode {
	return ce.ErrorCode
}

func (ce CustomError) Validations() []*ValidationError {
	return ce.ValidationErrors
}
