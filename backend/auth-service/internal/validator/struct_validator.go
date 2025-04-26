package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type StructValidatorError struct {
	Key      string
	UniqueId string
	Msg      string
}

type StructValidator struct {
	Struct interface{}
	Types  []interface{}
	Sl     validator.StructLevel
	v      *validator.Validate
	Errors []StructValidatorError
}

func NewStructValidator(data interface{}, types ...interface{}) StructValidator {
	sv := StructValidator{Struct: data, Types: []interface{}{data}, v: GetValidator()}
	return sv
}

func (sv *StructValidator) AddSimpleError(key string, msg string) StructValidator {
	sv.Errors = append(sv.Errors, StructValidatorError{
		Key:      key,
		UniqueId: uuid.New().String(),
		Msg:      msg,
	})
	return *sv
}

func (sv *StructValidator) Validate() error {
	sv.v.RegisterStructValidation(func(sl validator.StructLevel) {
		for _, err := range sv.Errors {
			RegisterTranslation(sv.v, err.UniqueId, err.Msg)
			sl.ReportError(err.Key, err.Key, err.Key, err.UniqueId, "")
		}
	}, sv.Types...)
	return sv.v.Struct(sv.Struct)
}
