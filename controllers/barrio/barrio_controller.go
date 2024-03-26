package barrioController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetBarrioDetalleById(c *gin.Context) {
	log.Debug("Sensor id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))
	var barrioDetailDto dto.BarrioDetailDto

	barrioDetailDto, err := service.BarrioService.GetBarrioDetalleById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, barrioDetailDto)
}
