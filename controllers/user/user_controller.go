package userController

import (
	"mvc-go/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetUserById(c *gin.Context) {
	log.Debug("User id: " + c.Param("id"))

	var userDto dto.UserDto
	userDto.Name = "Edu"
	userDto.LastName = "Gaite"

	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, "Users:")
}

func UserInsert(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, userDto)
}
