package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlightLeg_MarshalJSON(t *testing.T) {
	leg := &FlightLeg{
		Departure: "SFO",
		Arrival:   "ORD",
	}

	payload, err := json.Marshal(leg)
	assert.NoError(t, err)
	assert.Equal(t, string(payload), `["SFO","ORD"]`)
}

func TestFlightLeg_UnmarshalJSON(t *testing.T) {
	payload := `["SFO","ORD"]`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.NoError(t, err)
	assert.Equal(t, leg.Departure, AirportCode("SFO"))
	assert.Equal(t, leg.Arrival, AirportCode("ORD"))
}

func TestFlightLeg_UnmarshalJSON_ErrorInvalidJSON(t *testing.T) {
	payload := `{"SFO","ORD]`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.ErrorContains(t, err, "invalid character")
}

func TestFlightLeg_UnmarshalJSON_ErrorNotAJSONArray(t *testing.T) {
	payload := `{"departure": "SFO", "arrival": "ORD"}`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.ErrorContains(t, err, "cannot unmarshal object into Go value of type []interface {}")
}

func TestFlightLeg_UnmarshalJSON_ErrorMoreThan2AirportCodes(t *testing.T) {
	payload := `["SFO","ORD", "MIA"]`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.ErrorContains(t, err, "JSON array contains 3 entries")
}

func TestFlightLeg_UnmarshalJSON_ErrorLessThan2AirportCodes(t *testing.T) {
	payload := `["SFO"]`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.ErrorContains(t, err, "JSON array contains 1 entries")
}

func TestFlightLeg_UnmarshalJSON_ErrorEmptyArray(t *testing.T) {
	payload := `[]`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.ErrorContains(t, err, "JSON array contains 0 entries")
}

func TestFlightLeg_UnmarshalJSON_ErrorDepartureCodeIsNotAString(t *testing.T) {
	payload := `[5, "ORD"]`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.ErrorContains(t, err, "departure code is not a string")
}

func TestFlightLeg_UnmarshalJSON_ErrorArrivalCodeIsNotAString(t *testing.T) {
	payload := `["SFO", false]`

	var leg FlightLeg

	err := json.Unmarshal([]byte(payload), &leg)

	assert.ErrorContains(t, err, "arrival code is not a string")
}
