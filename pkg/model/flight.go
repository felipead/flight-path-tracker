package model

import (
	"encoding/json"
	"fmt"
)

type AirportCode string

type FlightLeg struct {
	Departure AirportCode
	Arrival   AirportCode
}

type FlightPath struct {
	Origin      AirportCode `json:"origin"`
	Destination AirportCode `json:"destination"`
	Legs        []FlightLeg `json:"legs"`
}

func (leg *FlightLeg) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unable to unmarshal flight leg: %w", err)
	}

	if len(v) != 2 {
		return fmt.Errorf("unable to unmarshal flight leg: JSON array contains %v entries", len(v))
	}

	var ok bool
	var origin, destination string

	if origin, ok = v[0].(string); !ok {
		return fmt.Errorf("unable to unmarshal flight leg: origin code is not a string")
	}
	if destination, ok = v[1].(string); !ok {
		return fmt.Errorf("unable to unmarshal flight leg: destination code is not a string")
	}

	leg.Departure = AirportCode(origin)
	leg.Arrival = AirportCode(destination)

	return nil
}

func (leg *FlightLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{
		string(leg.Departure),
		string(leg.Arrival),
	})
}
