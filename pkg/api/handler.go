package api

import (
	"github.com/gin-gonic/gin"

	"github.com/VolumeFi/flight-path-tracker/pkg/domain"
	"github.com/VolumeFi/flight-path-tracker/pkg/model"
)

func CalculateFlightPath(c *gin.Context) {
	var request CalculateFlightPathRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, NewErrorResponse(err))
		return
	}

	start, end, err := domain.FindFlightPathStartEnd(request.FlightLegs)
	if err != nil {
		c.AbortWithStatusJSON(400, NewErrorResponse(err))
		return
	}

	response := CalculateFlightPathResponse{
		FlightStartEnd: &model.FlightLeg{
			Origin:      start,
			Destination: end,
		},
	}

	c.JSON(200, response)
}
