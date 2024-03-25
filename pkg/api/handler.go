package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/felipead/flight-path-tracker/pkg/domain"
	"github.com/felipead/flight-path-tracker/pkg/validator"
)

var log = logrus.New()

func CalculateFlightPath(c *gin.Context) {
	var request CalculateFlightPathRequest

	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(400, NewErrorResponse(err))
		return
	}

	validate := validator.GetValidator()
	if err := validate.Struct(&request); err != nil {
		c.AbortWithStatusJSON(400, NewErrorResponse(err))
		return
	}

	log.WithFields(logrus.Fields{
		"FlightLegs": request.FlightLegs,
	}).Info("Calculating flight path")

	flightPath, err := domain.CalculateFlightPath(request.FlightLegs)
	if err != nil {
		c.AbortWithStatusJSON(400, NewErrorResponse(err))
		return
	}

	c.JSON(200, flightPath)
}
