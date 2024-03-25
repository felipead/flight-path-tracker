package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"

	"github.com/felipead/flight-path-tracker/pkg/model"
)

var singleton *validator.Validate

// InitValidator is not thread-safe and must be invoked before the server start serving requests.
func InitValidator() (err error) {
	if singleton == nil {
		singleton = validator.New()

		if err = model.RegisterAirportCodeValidation(singleton); err != nil {
			return
		}

		if err = singleton.RegisterValidation("notblank", validators.NotBlank); err != nil {
			return
		}
	}
	return
}

func GetValidator() *validator.Validate {
	return singleton
}
