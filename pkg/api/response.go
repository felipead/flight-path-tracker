package api

import (
	"github.com/felipead/flight-path-tracker/pkg/model"
)

type CalculateFlightPathResponse struct {
	FlightStartEnd *model.FlightLeg `json:"flight_start_end"`
}
