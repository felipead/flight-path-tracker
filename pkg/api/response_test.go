package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/VolumeFi/flight-path-tracker/pkg/model"
)

func TestMarshalCalculateFlightPathResponse(t *testing.T) {
	payload := CalculateFlightPathResponse{
		FlightStartEnd: &model.FlightLeg{
			Departure: "SFO",
			Arrival:   "EWR",
		},
	}

	jsonData, err := json.Marshal(payload)
	assert.NoError(t, err)
	assert.Equal(t, string(jsonData), "{\"flight_start_end\":[\"SFO\",\"EWR\"]}")
}
