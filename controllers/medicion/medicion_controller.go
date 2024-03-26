package medicionController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RegistrarMedicion(c *gin.Context) {
	var medicionDto dto.MedicionDto
	err := c.BindJSON(&medicionDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	medicionDto, er := service.MedicionService.RegistrarMedicion(medicionDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, medicionDto)
}
