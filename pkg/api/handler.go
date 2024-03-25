package api

import (
	"github.com/gin-gonic/gin"

	"github.com/felipead/flight-path-tracker/pkg/domain"
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

	c.JSON(200, flightPath)
}
