package app

import (
	productController "mvc-go/controllers/product"
	userController "mvc-go/controllers/user"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:product_id", productController.GetProductByEan)
	router.GET("/product", productController.GetProducts)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

}