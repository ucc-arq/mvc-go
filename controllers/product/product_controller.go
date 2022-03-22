package productController

import (
	"github.com/gin-gonic/gin"
	"net/http"	
)

func GetProductByEan(c *gin.Context) {
	c.JSON(http.StatusOK, "Buscando: " + c.Param("product_id"))
}

func GetProducts(c *gin.Context) {	
	c.JSON(http.StatusOK, "")
}
