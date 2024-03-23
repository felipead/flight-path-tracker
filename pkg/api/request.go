package api

import (
	"github.com/felipead/flight-path-tracker/pkg/model"
)

type CalculateFlightPathRequest struct {
	FlightLegs []model.FlightLeg `json:"flight_legs"`
}
