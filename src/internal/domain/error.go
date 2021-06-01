package domain

// CustomError - エラーコードを含めた構造体
type CustomError struct {
	ErrorCode        ErrorCode
	Value            error
	ValidationErrors []*ValidationError
}

// ValidationError - バリデーションエラー用構造体
type ValidationError struct {
	Field   string
	Message string
}

// ErrorCode - システムエラーの種類
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

// New - 指定したErrorCodeを持つCustomErrorを返す
func (ec ErrorCode) New(err error, ves ...*ValidationError) error {
	return CustomError{
		ErrorCode:        ec,
		Value:            err,
		ValidationErrors: ves,
	}
}

// Error - エラー内容を返す
func (ce CustomError) Error() string {
	return ce.Value.Error()
}

// Code - エラーコードを返す
func (ce CustomError) Code() ErrorCode {
	return ce.ErrorCode
}

// Validations - バリデーションエラーの詳細を返す
func (ce CustomError) Validations() []*ValidationError {
	return ce.ValidationErrors
}
