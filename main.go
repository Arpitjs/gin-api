package main

import (
	"gin-api/config"
	"gin-api/controller"
	"gin-api/dao"
	"gin-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	//db connection
	var (
		db                *gorm.DB                     = config.SetUpDB()
		dao               dao.ProductDAO               = dao.NewProductDAO(db)
		productService    services.ProductService      = services.NewProductService(dao)
		productController controller.ProductController = controller.NewproductController(productService)
	)

	//api routes
	router := gin.Default()

	router.GET("/products", productController.FindAll)

	router.POST("/product", productController.Create)

	router.PUT("/product/:id", productController.Update)

	router.DELETE("/product/:id", productController.Delete)

	router.Run()
}
