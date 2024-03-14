package api

import "github.com/VolumeFi/flight-path-tracker/pkg/model"

type CalculateFlightPathResponse struct {
	FlightOriginDestination *model.FlightLeg `json:"flight_origin_destination"`
}
