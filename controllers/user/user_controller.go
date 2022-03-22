package userController

import (
	"mvc-go/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {

	//c.Param("id")

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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, userDto)
}
