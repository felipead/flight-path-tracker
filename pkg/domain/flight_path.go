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

	path := NewPath[model.AirportCode]()

	for _, leg := range flightLegs {
		err := path.AddConnection(leg.Origin, leg.Destination)
		if err != nil {
			return "", "", fmt.Errorf("invalid flight path; %w", err)
		}
	}

	start, err := path.FindStart()
	if err != nil {
		return "", "", fmt.Errorf("invalid flight path; %w", err)
	}

	end, err := path.FindEnd()
	if err != nil {
		return "", "", fmt.Errorf("invalid flight path; %w", err)
	}

	return start, end, nil
}
