package api

import (
	"encoding/json"
	"github.com/felipead/flight-path-tracker/pkg/validator"
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

func TestValidateCalculateFlightPathRequest_ValidPayload(t *testing.T) {
	assert.NoError(t, validator.InitValidator())

	request := &CalculateFlightPathRequest{
		FlightLegs: []model.FlightLeg{
			{"ORD", "JFK"},
			{"SFO", "ORD"},
			{"JFK", "LHR"},
		},
	}

	validator := validator.GetValidator()
	assert.NoError(t, validator.Struct(request))
}

func TestValidateCalculateFlightPathRequest_EmptyList(t *testing.T) {
	assert.NoError(t, validator.InitValidator())

	request := &CalculateFlightPathRequest{
		FlightLegs: []model.FlightLeg{},
	}

	validator := validator.GetValidator()
	assert.ErrorContains(
		t, validator.Struct(request),
		"validation for 'FlightLegs' failed on the 'notblank' tag",
	)
}

func TestValidateCalculateFlightPathRequest_EmptyDepartureAirportCode(t *testing.T) {
	assert.NoError(t, validator.InitValidator())

	request := &CalculateFlightPathRequest{
		FlightLegs: []model.FlightLeg{
			{"ORD", "JFK"},
			{"", "ORD"},
			{"JFK", "LHR"},
		},
	}

	validator := validator.GetValidator()
	assert.ErrorContains(
		t, validator.Struct(request),
		"Error:Field validation for 'Departure' failed on the 'required' tag",
	)
}

func TestValidateCalculateFlightPathRequest_InvalidDepartureAirportCode(t *testing.T) {
	assert.NoError(t, validator.InitValidator())

	request := &CalculateFlightPathRequest{
		FlightLegs: []model.FlightLeg{
			{"ORD", "JFK"},
			{"SSFF5", "ORD"},
			{"JFK", "LHR"},
		},
	}

	validator := validator.GetValidator()
	assert.ErrorContains(
		t, validator.Struct(request),
		"Error:Field validation for 'Departure' failed on the 'airport_code' tag",
	)
}

func TestValidateCalculateFlightPathRequest_EmptyArrivalAirportCode(t *testing.T) {
	assert.NoError(t, validator.InitValidator())

	request := &CalculateFlightPathRequest{
		FlightLegs: []model.FlightLeg{
			{"ORD", "JFK"},
			{"SFO", ""},
			{"JFK", "LHR"},
		},
	}

	validator := validator.GetValidator()
	assert.ErrorContains(
		t, validator.Struct(request),
		"Error:Field validation for 'Arrival' failed on the 'required' tag",
	)
}

func TestValidateCalculateFlightPathRequest_InvalidArrivalAirportCode(t *testing.T) {
	assert.NoError(t, validator.InitValidator())

	request := &CalculateFlightPathRequest{
		FlightLegs: []model.FlightLeg{
			{"ORD", "JFK"},
			{"SFO", "555"},
			{"JFK", "LHR"},
		},
	}

	validator := validator.GetValidator()
	assert.ErrorContains(
		t, validator.Struct(request),
		"Error:Field validation for 'Arrival' failed on the 'airport_code' tag",
	)
}
