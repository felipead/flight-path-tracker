package api

import (
	"github.com/VolumeFi/flight-path-tracker/pkg/model"
)

type CalculateFlightPathResponse struct {
	FlightStartEnd *model.FlightLeg `json:"flight_start_end"`
}
