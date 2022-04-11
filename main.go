package main

import (
	"gin-api/controller"
	"gin-api/dao"
	"gin-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price string
}

func main() {
	//db connection

	dsn := "host=localhost user=arpit password=leapfrog123 dbname=crudapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	var (
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
