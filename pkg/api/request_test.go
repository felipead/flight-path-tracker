package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/felipead/flight-path-tracker/pkg/model"
)

func TestUnmarshalCalculateFlightPathRequest_ValidPayload(t *testing.T) {
	payload := `{
	"flight_legs": [
		["IND", "EWR"],
		["SFO", "ATL"],
		["GSO", "IND"],
		["ATL", "GSO"]
	]
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.NoError(t, err)

	assert.Equal(t, len(request.FlightLegs), 4)
	assert.Equal(t, request.FlightLegs, []model.FlightLeg{
		{
			Departure: "IND",
			Arrival:   "EWR",
		},
		{
			Departure: "SFO",
			Arrival:   "ATL",
		},
		{
			Departure: "GSO",
			Arrival:   "IND",
		},
		{
			Departure: "ATL",
			Arrival:   "GSO",
		},
	})
}

func TestUnmarshalCalculateFlightPathRequest_EmptyList(t *testing.T) {
	payload := `{
	"flight_legs": []
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.NoError(t, err)

	assert.Equal(t, len(request.FlightLegs), 0)
}

func TestUnmarshalCalculateFlightPathRequest_NotAList(t *testing.T) {
	payload := `{
	"flight_legs": 333
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.ErrorContains(t, err, "cannot unmarshal number into Go struct field")
}

func TestUnmarshalCalculateFlightPathRequest_EmptyJSON(t *testing.T) {
	payload := `{
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.NoError(t, err)

	assert.Equal(t, len(request.FlightLegs), 0)
}

func TestUnmarshalCalculateFlightPathRequest_InvalidFlight_1Entry(t *testing.T) {
	payload := `{
	"flight_legs": [
		["IND", "EWR"],
		["SFO"],
		["GSO", "IND"],
		["ATL", "GSO"]
	]
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.EqualError(t, err, "unable to unmarshal flight leg: JSON array contains 1 entries")
}

func TestUnmarshalCalculateFlightPathRequest_InvalidFlight_3Entries(t *testing.T) {
	payload := `{
	"flight_legs": [
		["IND", "EWR"],
		["SFO", "ATL", "FOO"],
		["GSO", "IND"],
		["ATL", "GSO"]
	]
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.EqualError(t, err, "unable to unmarshal flight leg: JSON array contains 3 entries")
}

func TestUnmarshalCalculateFlightPathRequest_InvalidFlight_OriginCodeIsNotAString(t *testing.T) {
	payload := `{
	"flight_legs": [
		["IND", "EWR"],
		[555, "ATL"],
		["GSO", "IND"],
		["ATL", "GSO"]
	]
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.EqualError(t, err, "unable to unmarshal flight leg: departure code is not a string")
}

func TestUnmarshalCalculateFlightPathRequest_InvalidFlight_DestinationCodeIsNotAString(t *testing.T) {
	payload := `{
	"flight_legs": [
		["IND", "EWR"],
		["SFO", 333],
		["GSO", "IND"],
		["ATL", "GSO"]
	]
}`
	var request CalculateFlightPathRequest
	err := json.Unmarshal([]byte(payload), &request)
	assert.EqualError(t, err, "unable to unmarshal flight leg: arrival code is not a string")
}
