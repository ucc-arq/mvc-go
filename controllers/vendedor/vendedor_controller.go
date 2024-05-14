package vendedorController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetVendedorById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var vendedorDto dto.VendedorDto

	vendedorDto, err := service.VendedorService.GetVendedorById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, vendedorDto)
}
