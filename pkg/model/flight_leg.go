package model

import (
	"encoding/json"
	"fmt"
)

type FlightLeg struct {
	Departure AirportCode `validate:"required,airport_code"`
	Arrival   AirportCode `validate:"required,airport_code"`
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
	var departure, arrival string

	if departure, ok = v[0].(string); !ok {
		return fmt.Errorf("unable to unmarshal flight leg: departure code is not a string")
	}
	if arrival, ok = v[1].(string); !ok {
		return fmt.Errorf("unable to unmarshal flight leg: arrival code is not a string")
	}

	leg.Departure = AirportCode(departure)
	leg.Arrival = AirportCode(arrival)

	return nil
}

func (leg *FlightLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{
		string(leg.Departure),
		string(leg.Arrival),
	})
}
