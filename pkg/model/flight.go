package model

import (
	"encoding/json"
	"fmt"
)

type FlightLeg struct {
	// OriginCode is the origin IATA airport code
	OriginCode string
	// DestinationCode is the destination IATA airport code
	DestinationCode string
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
	var originCode, destinationCode string

	if originCode, ok = v[0].(string); !ok {
		return fmt.Errorf("unable to unmarshal flight leg: origin code is not a string")
	}
	if destinationCode, ok = v[1].(string); !ok {
		return fmt.Errorf("unable to unmarshal flight leg: destination code is not a string")
	}

	leg.OriginCode = originCode
	leg.DestinationCode = destinationCode

	return nil
}

func (leg *FlightLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string{
		leg.OriginCode,
		leg.DestinationCode,
	})
}
