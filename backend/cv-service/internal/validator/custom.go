package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// Register single validator & it's translation
func Register(validate *validator.Validate, key string, fn func(fl FieldLevel) bool, errMsg string) {
	validate.RegisterValidation(key, func(vFl validator.FieldLevel) bool { return fn(vFl) })
	RegisterTranslation(validate, key, errMsg)
}

// Register Translation
func RegisterTranslation(validate *validator.Validate, key string, errMsg string) {
	validate.RegisterTranslation(key, trans, func(ut ut.Translator) error {
		return ut.Add(key, errMsg, true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(key, fe.Field())

		return t
	})
}
