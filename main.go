package main

import (
	"github.com/itsmemz53/go-sensors-read/controllers"
	"github.com/itsmemz53/go-sensors-read/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/sensors", controllers.GetAllSensors)
	r.GET("/sensors/:id", controllers.GetSensorData)
	r.POST("/sensor", controllers.PostSensor)
	r.POST("/sensor-reading", controllers.PostSensorReadings)

	// Run the server
	r.Run()
}
