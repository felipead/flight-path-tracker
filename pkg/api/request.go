package api

import "github.com/VolumeFi/flight-path-tracker/pkg/model"

type CalculateFlightPathRequest struct {
	FlightLegs []model.FlightLeg `json:"flight_legs"`
}
