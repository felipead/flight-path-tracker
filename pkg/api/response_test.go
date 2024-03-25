package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/felipead/flight-path-tracker/pkg/model"
)

func TestMarshalCalculateFlightPathResponse(t *testing.T) {
	payload := CalculateFlightPathResponse{
		FlightPath: &model.FlightPath{
			Origin:      "SFO",
			Destination: "EWR",
			Legs: []model.FlightLeg{
				{"SFO", "ATL"},
				{"ATL", "GSO"},
				{"GSO", "IND"},
				{"IND", "EWR"},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	assert.NoError(t, err)
	assert.Equal(t, string(jsonData),
		`{"flight_path":{"origin":"SFO","destination":"EWR",`+
			`"legs":[["SFO","ATL"],["ATL","GSO"],["GSO","IND"],["IND","EWR"]]}}`,
	)
}
