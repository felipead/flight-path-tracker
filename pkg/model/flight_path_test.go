package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlightPath_MarshalJSON(t *testing.T) {
	payload := &FlightPath{
		Origin:      "SFO",
		Destination: "EWR",
		FlightLegs: []FlightLeg{
			{"SFO", "ATL"},
			{"ATL", "GSO"},
			{"GSO", "IND"},
			{"IND", "EWR"},
		},
	}

	jsonData, err := json.Marshal(payload)
	assert.NoError(t, err)
	assert.Equal(t, string(jsonData),
		`{"origin":"SFO","destination":"EWR",`+
			`"flight_legs":[["SFO","ATL"],["ATL","GSO"],["GSO","IND"],["IND","EWR"]]}`,
	)
}
