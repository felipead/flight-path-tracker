package domain

import (
	"testing"

	"github.com/VolumeFi/flight-path-tracker/pkg/model"
)

func TestFindFlightPathStartEnd_Success(t *testing.T) {
	tests := []struct {
		name       string
		flightLegs []model.FlightLeg
		wantStart  model.AirportCode
		wantEnd    model.AirportCode
	}{
		{
			name: "given one flight leg",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "SFO",
					Destination: "CNF",
				},
			},
			wantStart: "SFO",
			wantEnd:   "CNF",
		},
		{
			name: "given two flight legs, sorted",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "SFO",
					Destination: "CNF",
				},
				{
					Origin:      "CNF",
					Destination: "MIA",
				},
			},
			wantStart: "SFO",
			wantEnd:   "MIA",
		},
		{
			name: "given two flight legs, unsorted",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "ATL",
					Destination: "EWR",
				},
				{
					Origin:      "SFO",
					Destination: "ATL",
				},
			},
			wantStart: "SFO",
			wantEnd:   "EWR",
		},
		{
			name: "given a few unsorted flight legs, sample 1",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "IND",
					Destination: "EWR",
				},
				{
					Origin:      "SFO",
					Destination: "ATL",
				},
				{
					Origin:      "GSO",
					Destination: "IND",
				},
				{
					Origin:      "ATL",
					Destination: "GSO",
				},
			},
			wantStart: "SFO",
			wantEnd:   "EWR",
		},
		{
			name: "given a few unsorted flight legs, sample 2",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
			},
			wantStart: "CNF",
			wantEnd:   "LHR",
		},
		{
			name: "given a few sorted flight legs",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
			},
			wantStart: "CNF",
			wantEnd:   "LHR",
		},
		{
			name: "given a few reversely sorted flight legs",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
			},
			wantStart: "CNF",
			wantEnd:   "LHR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd, err := FindFlightPathStartEnd(tt.flightLegs)
			if err != nil {
				t.Errorf("FindFlightPathStartEnd() error = %v", err)
				return
			}
			if gotStart != tt.wantStart {
				t.Errorf("FindFlightPathStartEnd() gotStart = %v, wantStart = %v", gotStart, tt.wantStart)
			}
			if gotEnd != tt.wantEnd {
				t.Errorf("FindFlightPathStartEnd() gotEnd = %v, wantEnd = %v", gotEnd, tt.wantEnd)
			}
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
					Origin:      "SFO",
					Destination: "SFO",
				},
			},
			wantError: "invalid flight path - a flight leg can't point to itself",
		},
		{
			name: "when the flight path has an airport with more than one outbound legs (repeats)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
			},
			wantError: "invalid flight path - found more than one outbound airports for ORD",
		},
		{
			name: "when the flight path has an airport with more than one outbound legs (branch)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "ORD",
					Destination: "MIA",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
			},
			wantError: "invalid flight path - found more than one outbound airports for ORD",
		},
		{
			name: "when the flight path has an airport with more than one outbound legs (loop)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "SFO",
					Destination: "ORD",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
			},
			wantError: "invalid flight path - found more than one outbound airports for SFO",
		},
		{

			name: "when the flight path has an airport with more than one inbound legs (branch)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "HND",
					Destination: "SFO",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
			},
			wantError: "invalid flight path - found more than one inbound airports for SFO",
		},
		{
			name: "when the flight path has an airport with more than one inbound legs (loop)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "LHR",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
			},
			wantError: "invalid flight path - found more than one inbound airports for JFK",
		},
		{
			name: "when there's no start or end (loop)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
				{
					Origin:      "LHR",
					Destination: "CNF",
				},
			},
			wantError: "invalid flight path - loop detected",
		},
		{
			name: "when there's no start but there's an end (loop)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
				{
					Origin:      "YUL",
					Destination: "CNF",
				},
			},
			wantError: "invalid flight path - found more than one outbound airports for YUL",
		},
		{
			name: "when there's no end but there's an start (loop)",
			flightLegs: []model.FlightLeg{
				{
					Origin:      "CNF",
					Destination: "GRU",
				},
				{
					Origin:      "GRU",
					Destination: "MIA",
				},
				{
					Origin:      "MIA",
					Destination: "ORD",
				},
				{
					Origin:      "ORD",
					Destination: "SFO",
				},
				{
					Origin:      "SFO",
					Destination: "YUL",
				},
				{
					Origin:      "YUL",
					Destination: "JFK",
				},
				{
					Origin:      "JFK",
					Destination: "LHR",
				},
				{
					Origin:      "LHR",
					Destination: "YUL",
				},
			},
			wantError: "invalid flight path - found more than one inbound airports for YUL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, gotError := FindFlightPathStartEnd(tt.flightLegs)
			if gotError == nil {
				t.Errorf("FindFlightPathStartEnd() did not fail, but an error was expected; start = %v, end = %v", start, end)
				return
			}
			if gotError.Error() != tt.wantError {
				t.Errorf("FindFlightPathStartEnd() gotError = %v, wantError = %v", gotError, tt.wantError)
			}
		})
	}
}
