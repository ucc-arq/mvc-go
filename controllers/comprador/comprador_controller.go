package compradorController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCompradorById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var compradorDto dto.CompradorDto

	compradorDto, err := service.CompradorService.GetCompradorById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, compradorDto)
}
