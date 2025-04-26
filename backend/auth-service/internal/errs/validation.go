package errs

import (
	"auth-service/internal/validator"

	go_validator "github.com/go-playground/validator/v10"
)

// MissingField is an error type that can be used when
// validating input fields that do not have a value, but should
type MissingField string

func (e MissingField) Error() string {
	return string(e) + " is required"
}

// InputUnwanted is an error type that can be used when
// validating input fields that have a value, but should should not
type InputUnwanted string

func (e InputUnwanted) Error() string {
	return string(e) + " has a value, but should be nil"
}

// --------------------

type ValidationErrorData map[string]([]string)

func NewValidationErrorMessages(err go_validator.ValidationErrors) ValidationErrorData {
	data := ValidationErrorData{}

	for _, validationError := range err {
		data.AppendValidationError(validationError)
	}

	return data

}

func (data *ValidationErrorData) AppendValidationError(fieldError go_validator.FieldError) {
	tag := fieldError.Field()
	msg := fieldError.Translate(validator.GetTranslator())
	(*data)[tag] = append((*data)[tag], msg)
}
