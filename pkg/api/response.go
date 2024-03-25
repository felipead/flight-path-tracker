package api

import (
	"github.com/felipead/flight-path-tracker/pkg/model"
)

type CalculateFlightPathResponse struct {
	FlightPath *model.FlightPath `json:"flight_path"`
}
