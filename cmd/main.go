package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/felipead/flight-path-tracker/pkg/api"
)

func main() {
	router := gin.Default()

	router.POST("/calculate", api.CalculateFlightPath)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
