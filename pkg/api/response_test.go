package api

import (
	"encoding/json"
	"github.com/VolumeFi/flight-path-tracker/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalCalculateFlightPathResponse(t *testing.T) {
	payload := CalculateFlightPathResponse{
		FlightSummary: &model.FlightLeg{
			OriginCode:      "SFO",
			DestinationCode: "EWR",
		},
	}

	jsonData, err := json.Marshal(payload)
	assert.NoError(t, err)
	assert.Equal(t, string(jsonData), "{\"flight_summary\":[\"SFO\",\"EWR\"]}")
}
