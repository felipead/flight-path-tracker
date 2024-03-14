package api

import (
	"github.com/gin-gonic/gin"
)

func CalculateFlightPath(c *gin.Context) {
	var request CalculateFlightPathRequest
	if err := c.BindJSON(&request); err != nil {
		_ = c.Error(err)
		return
	}
}
