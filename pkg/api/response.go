package api

import "github.com/VolumeFi/flight-path-tracker/pkg/model"

type CalculateFlightPathResponse struct {
	FlightSummary *model.FlightLeg `json:"flight_summary"`
}
