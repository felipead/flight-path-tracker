package api

import (
	"encoding/json"
	"github.com/VolumeFi/flight-path-tracker/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalCalculateFlightPathResponse(t *testing.T) {
	payload := CalculateFlightPathResponse{
		FlightStartEnd: &model.FlightLeg{
			Origin:      "SFO",
			Destination: "EWR",
		},
	}

	jsonData, err := json.Marshal(payload)
	assert.NoError(t, err)
	assert.Equal(t, string(jsonData), "{\"flight_start_end\":[\"SFO\",\"EWR\"]}")
}
