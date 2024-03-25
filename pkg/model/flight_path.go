package model

type FlightPath struct {
	Origin      AirportCode `json:"origin"`
	Destination AirportCode `json:"destination"`
	FlightLegs  []FlightLeg `json:"flight_legs"`
}
