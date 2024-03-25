package domain

import (
	"errors"
	"fmt"

	"github.com/felipead/flight-path-tracker/pkg/model"
)

func CalculateFlightPath(flightLegs []model.FlightLeg) (*model.FlightPath, error) {
	if len(flightLegs) == 0 {
		return nil, errors.New("empty flight path")
	}

	path := NewPath[model.AirportCode]()

	for _, leg := range flightLegs {
		err := path.AddConnection(leg.Departure, leg.Arrival)
		if err != nil {
			return nil, fmt.Errorf("invalid flight path; %w", err)
		}
	}

	start, err := path.FindStart()
	if err != nil {
		return nil, fmt.Errorf("invalid flight path; %w", err)
	}

	end, err := path.FindEnd()
	if err != nil {
		return nil, fmt.Errorf("invalid flight path; %w", err)
	}

	sortedLegs, err := sortFlightLegs(path, start, end)
	if err != nil {
		return nil, err
	}

	return &model.FlightPath{
		Origin:      start,
		Destination: end,
		FlightLegs:  sortedLegs,
	}, nil
}

func sortFlightLegs(path *Path[model.AirportCode], start, end model.AirportCode) ([]model.FlightLeg, error) {
	sortedLegs := make([]model.FlightLeg, 0, path.Length())

	//
	// Since Path is *guaranteed* to not have branches or loops, we don't need to worry about the loop below
	// never ending. We need to worry though if the path is disconnected or partitioned, which is captured in the
	// check below.
	//

	this := start
	for this != end {
		next := path.GetNext(this)

		if next == "" {
			return nil, fmt.Errorf("disconnected flight path; there's no flight leg leaving airport %v", this)
		}

		sortedLegs = append(sortedLegs, model.FlightLeg{
			Departure: this,
			Arrival:   next,
		})
		this = next
	}

	return sortedLegs, nil
}
