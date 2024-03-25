package api

import (
	"github.com/felipead/flight-path-tracker/pkg/validator"
)

// Init is supposed to be called before the server starts serving API requests
func Init() (err error) {
	err = validator.InitValidator()
	return
}
