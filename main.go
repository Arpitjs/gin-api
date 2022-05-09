package main

import (
	"gin-api/config"
	"gin-api/controller"
	"gin-api/dao"

	"gin-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)
//
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// @title Swagger  demo service API
// @version 1.0
// @description This is demo server.
// @termsOfService demo.com

// @contact.name API Support
// @contact.url http://demo.com/support

// @host localhost:8091
// @BasePath /api/v1

// @in header
// @name Authorization

func main() {

	var (
		db                *gorm.DB                     = config.SetUpDB()
		dao               dao.ProductDAO               = dao.NewProductDAO(db)
		productService    services.ProductService      = services.NewProductService(dao)
		productController controller.ProductController = controller.NewproductController(productService)
	)

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/products", productController.FindAll)

	router.POST("/products", productController.Create)

	router.PUT("/product/:id", productController.Update)

	router.DELETE("/product/:id", productController.Delete)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
