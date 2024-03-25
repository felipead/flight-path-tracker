package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/felipead/flight-path-tracker/pkg/api"
)

func main() {
	if err := api.Init(); err != nil {
		log.Panic(err)
	}

	router := gin.Default()
	router.POST("/calculate", api.CalculateFlightPath)

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
