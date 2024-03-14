package domain

import (
	"errors"
	"fmt"

	"github.com/VolumeFi/flight-path-tracker/pkg/model"
)

func FindFlightPathStartEnd(flightLegs []model.FlightLeg) (model.AirportCode, model.AirportCode, error) {
	if len(flightLegs) == 0 {
		return "", "", errors.New("empty flight path")
	}

	airports := make(map[model.AirportCode]bool)
	for _, e := range flightLegs {
		if e.Origin == e.Destination {
			return "", "", errors.New("invalid flight path - a flight leg can't point to itself")
		}

		airports[e.Origin] = true
		airports[e.Destination] = true
	}

	outboundOf := make(map[model.AirportCode]model.AirportCode)
	for v := range airports {
		for _, e := range flightLegs {
			if e.Origin == v {
				if outboundOf[v] != "" {
					return "", "", fmt.Errorf(
						"invalid flight path - found more than one outbound airports for %v", v)
				}
				outboundOf[v] = e.Destination
			}
		}
	}

	inboundOf := make(map[model.AirportCode]model.AirportCode)
	for v := range airports {
		for _, e := range flightLegs {
			if e.Destination == v {
				if inboundOf[v] != "" {
					return "", "", fmt.Errorf(
						"invalid flight path - found more than one inbound airports for %v", v)
				}
				inboundOf[v] = e.Origin
			}
		}
	}

	var start, end model.AirportCode
	for v := range airports {
		if inboundOf[v] == "" {
			start = v
			continue
		}
		if outboundOf[v] == "" {
			end = v
			continue
		}
	}

	if start == "" || end == "" {
		return "", "", errors.New("invalid flight path - loop detected")
	}

	return start, end, nil
}
