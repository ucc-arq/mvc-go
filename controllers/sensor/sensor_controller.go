package sensorController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetSensorById(c *gin.Context) {
	log.Debug("Sensor id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var sensorDetailDto dto.SensorDetailDto

	sensorDetailDto, err := service.SensorService.GetSensorById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, sensorDetailDto)
}

func SensorInsert(c *gin.Context) {
	var sensorDto dto.SensorDto
	err := c.BindJSON(&sensorDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	sensorDto, er := service.SensorService.InsertSensor(sensorDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, sensorDto)
}

func ActivarSensor(c *gin.Context) {
	log.Debug("Sensor id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var sensorDto dto.SensorDto
	sensorDto, er := service.SensorService.ActivarSensor(id)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, sensorDto)
}

func PausarSensor(c *gin.Context) {
	log.Debug("Sensor id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var sensorDto dto.SensorDto
	sensorDto, er := service.SensorService.PausarSensor(id)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, sensorDto)
}
