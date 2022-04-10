package controller

import (
	"gin-api/entity"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll(context *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewproductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) FindAll(context *gin.Context) {
	var products []entity.Product = c.productService.FindAll()
	context.JSON(200, products)
}
