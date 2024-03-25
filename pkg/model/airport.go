package model

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

type AirportCode string

func (code AirportCode) IsValid() bool {
	return len(code) == 3 &&
		unicode.IsLetter(rune(code[0])) &&
		unicode.IsLetter(rune(code[1])) &&
		unicode.IsLetter(rune(code[2]))
}

func RegisterAirportCodeValidation(validate *validator.Validate) error {
	return validate.RegisterValidation("airport_code", func(f validator.FieldLevel) bool {
		value := f.Field().Interface().(AirportCode)
		return value.IsValid()
	})
}
