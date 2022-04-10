package main

import (
	"gin-api/controller"
	"gin-api/dao"
	"gin-api/services"
	"net/http"

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

	router.POST("/product", func(c *gin.Context) {
		code := c.PostForm("Code")
		price := c.PostForm("Price")
		db.Create(&Product{Code: code, Price: price})
		c.String(http.StatusOK, "Product created!")
	})

	router.PUT("/product/:id", func(c *gin.Context) {
		var product Product
		id := c.Param("id")
		price := c.PostForm("Price")
		db.Model(&product).Where("id", id).Update("Price", price)
		c.String(http.StatusOK, "Product updated!")
	})

	router.DELETE("/product/:id", func(c *gin.Context) {
		var product Product
		id := c.Param("id")
		db.Model(&product).Where("id", id).Delete(&product)
		c.String(http.StatusOK, "Product deleted!")
	})

	router.Run()
}
