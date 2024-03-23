package api

import (
	"github.com/gin-gonic/gin"

	"github.com/felipead/flight-path-tracker/pkg/domain"
	"github.com/felipead/flight-path-tracker/pkg/model"
)

func CalculateFlightPath(c *gin.Context) {
	var request CalculateFlightPathRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, NewErrorResponse(err))
		return
	}

	flightPath, err := domain.CalculateFlightPath(request.FlightLegs)
	if err != nil {
		c.AbortWithStatusJSON(400, NewErrorResponse(err))
		return
	}

	response := CalculateFlightPathResponse{
		FlightStartEnd: &model.FlightLeg{
			Departure: flightPath.Origin,
			Arrival:   flightPath.Destination,
		},
	}

	c.JSON(200, response)
}
