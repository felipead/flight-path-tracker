package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/felipead/flight-path-tracker/pkg/model"
)

func TestCalculateFlightPath_Success(t *testing.T) {
	tests := []struct {
		name            string
		flightLegs      []model.FlightLeg
		wantOrigin      model.AirportCode
		wantDestination model.AirportCode
		wantSortedLegs  []model.FlightLeg
	}{
		{
			name: "given one flight leg",
			flightLegs: []model.FlightLeg{
				{
					Departure: "SFO",
					Arrival:   "CNF",
				},
			},
			wantOrigin:      "SFO",
			wantDestination: "CNF",
			wantSortedLegs: []model.FlightLeg{
				{"SFO", "CNF"},
			},
		},
		{
			name: "given two flight legs, sorted",
			flightLegs: []model.FlightLeg{
				{
					Departure: "SFO",
					Arrival:   "CNF",
				},
				{
					Departure: "CNF",
					Arrival:   "MIA",
				},
			},
			wantOrigin:      "SFO",
			wantDestination: "MIA",
			wantSortedLegs: []model.FlightLeg{
				{"SFO", "CNF"},
				{"CNF", "MIA"},
			},
		},
		{
			name: "given two flight legs, unsorted",
			flightLegs: []model.FlightLeg{
				{
					Departure: "ATL",
					Arrival:   "EWR",
				},
				{
					Departure: "SFO",
					Arrival:   "ATL",
				},
			},
			wantOrigin:      "SFO",
			wantDestination: "EWR",
			wantSortedLegs: []model.FlightLeg{
				{"SFO", "ATL"},
				{"ATL", "EWR"},
			},
		},
		{
			name: "given a few unsorted flight legs, sample 1",
			flightLegs: []model.FlightLeg{
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
			},
			wantOrigin:      "SFO",
			wantDestination: "EWR",
			wantSortedLegs: []model.FlightLeg{
				{"SFO", "ATL"},
				{"ATL", "GSO"},
				{"GSO", "IND"},
				{"IND", "EWR"},
			},
		},
		{
			name: "given a few unsorted flight legs, sample 2",
			flightLegs: []model.FlightLeg{
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
			},
			wantOrigin:      "CNF",
			wantDestination: "LHR",
			wantSortedLegs: []model.FlightLeg{
				{"CNF", "GRU"},
				{"GRU", "MIA"},
				{"MIA", "ORD"},
				{"ORD", "SFO"},
				{"SFO", "YUL"},
				{"YUL", "JFK"},
				{"JFK", "LHR"},
			},
		},
		{
			name: "given a few sorted flight legs",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
			},
			wantOrigin:      "CNF",
			wantDestination: "LHR",
			wantSortedLegs: []model.FlightLeg{
				{"CNF", "GRU"},
				{"GRU", "MIA"},
				{"MIA", "ORD"},
				{"ORD", "SFO"},
				{"SFO", "YUL"},
				{"YUL", "JFK"},
				{"JFK", "LHR"},
			},
		},
		{
			name: "given a few reversely sorted flight legs",
			flightLegs: []model.FlightLeg{
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
			},
			wantOrigin:      "CNF",
			wantDestination: "LHR",
			wantSortedLegs: []model.FlightLeg{
				{"CNF", "GRU"},
				{"GRU", "MIA"},
				{"MIA", "ORD"},
				{"ORD", "SFO"},
				{"SFO", "YUL"},
				{"YUL", "JFK"},
				{"JFK", "LHR"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateFlightPath(tt.flightLegs)
			assert.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, got.Origin, tt.wantOrigin)
			assert.Equal(t, got.Destination, tt.wantDestination)
			assert.Equal(t, got.FlightLegs, tt.wantSortedLegs)
		})
	}
}

func TestFindFlightPathStartEnd_Error(t *testing.T) {
	tests := []struct {
		name       string
		flightLegs []model.FlightLeg
		wantError  string
	}{
		{
			name:       "when given an empty flight path",
			flightLegs: []model.FlightLeg{},
			wantError:  "empty flight path",
		},
		{
			name: "when given a flight leg pointing to itself",
			flightLegs: []model.FlightLeg{
				{
					Departure: "SFO",
					Arrival:   "SFO",
				},
			},
			wantError: `invalid flight path; invalid connection - "from" and "to" are the same`,
		},
		{
			name: "when the flight path has an airport with more than one outbound legs (repeats)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
			},
			wantError: `invalid flight path; invalid connection - "from" already has an outbound connection`,
		},
		{
			name: "when the flight path has an airport with more than one outbound legs (branch)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "ORD",
					Arrival:   "MIA",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
			},
			wantError: `invalid flight path; invalid connection - "from" already has an outbound connection`,
		},
		{
			name: "when the flight path has an airport with more than one outbound legs (loop)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "SFO",
					Arrival:   "ORD",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
			},
			wantError: `invalid flight path; invalid connection - "from" already has an outbound connection`,
		},
		{

			name: "when the flight path has an airport with more than one inbound legs (branch)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "HND",
					Arrival:   "SFO",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
			},
			wantError: `invalid flight path; invalid connection - "to" already has an inbound connection`,
		},
		{
			name: "when the flight path has an airport with more than one inbound legs (loop)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "LHR",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
			},
			wantError: `invalid flight path; invalid connection - "to" already has an inbound connection`,
		},
		{
			name: "when there's no start or end (loop)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
				{
					Departure: "LHR",
					Arrival:   "CNF",
				},
			},
			wantError: `invalid flight path; unable to find start of path - there's a loop`,
		},
		{
			name: "when there's no start but there's an end (loop)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
				{
					Departure: "YUL",
					Arrival:   "CNF",
				},
			},
			wantError: `invalid flight path; invalid connection - "from" already has an outbound connection`,
		},
		{
			name: "when there's no end but there's an start (loop)",
			flightLegs: []model.FlightLeg{
				{
					Departure: "CNF",
					Arrival:   "GRU",
				},
				{
					Departure: "GRU",
					Arrival:   "MIA",
				},
				{
					Departure: "MIA",
					Arrival:   "ORD",
				},
				{
					Departure: "ORD",
					Arrival:   "SFO",
				},
				{
					Departure: "SFO",
					Arrival:   "YUL",
				},
				{
					Departure: "YUL",
					Arrival:   "JFK",
				},
				{
					Departure: "JFK",
					Arrival:   "LHR",
				},
				{
					Departure: "LHR",
					Arrival:   "YUL",
				},
			},
			wantError: `invalid flight path; invalid connection - "to" already has an inbound connection`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flightPath, gotError := CalculateFlightPath(tt.flightLegs)
			if gotError == nil {
				t.Errorf("CalculateFlightPath() did not fail, but an error was expected; path = %v", flightPath)
				return
			}
			if gotError.Error() != tt.wantError {
				t.Errorf("CalculateFlightPath() gotError = %v, wantError = %v", gotError, tt.wantError)
			}
		})
	}
}
