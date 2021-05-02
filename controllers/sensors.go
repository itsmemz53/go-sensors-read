
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsmemz53/go-sensors-read/models"
)

type PostReadingBody struct {
	RequestId  string `json:"request_id" binding:"required"`
	Data string `json:"data" binding:"required"`
	SensorId int64 `json:"sensor_id" binding:"required"`
}

type PostSensorBody struct {
	Name  string `json:"name" binding:"required"`
	MacAddress string `json:"mac_address" binding:"required"`
}

// GET /sensors
// Get all sensors
func GetAllSensors(c *gin.Context) {
	var sensors []models.Sensors
	models.DB.Find(&sensors)

	c.JSON(http.StatusOK, gin.H{"data": sensors})
}

// GET /sensors/:id
// Get sensor data
func GetSensorData(c *gin.Context) {
	// Get model if exist
	var readings []models.Readings
	if err := models.DB.Where("sensor_id = ?", c.Param("id")).Find(&readings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": readings})
}

// POST /sensor
// Create new sensor
func PostSensor(c *gin.Context) {
	// Validate input
	var input PostSensorBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var sensorExist models.Sensors
	if err := models.DB.Where("mac_address = ?", input.MacAddress).First(&sensorExist).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sensor already exist with same mac address!"})
		return
	}

	// Create sensor
	sensor := models.Sensors{Name: input.Name, MacAddress: input.MacAddress}
	models.DB.Create(&sensor)

	c.JSON(http.StatusOK, gin.H{"data": sensor})
}

// POST /sensor-reading
// Post Sensors Reading
func PostSensorReadings(c *gin.Context) {
	// Validate input
	var input PostReadingBody
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check whether sensor exist
	var sensorExist models.Sensors
	if err := models.DB.Where("id = ?", input.SensorId).First(&sensorExist).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sensor does not exist"})
		return
	}

	// Check duplicate readings
	var readingExist models.Readings
	if err := models.DB.Where("request_id = ?", input.RequestId).First(&readingExist).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reading already exist with same request id"})
		return
	}

	// Create sensor
	reading := models.Readings{RequestId: input.RequestId, Data: input.Data, Sensors: &sensorExist}
	models.DB.Create(&reading)

	c.JSON(http.StatusOK, gin.H{"data": reading})
}
