package controller

import (
	"gin-api/entity"
	"gin-api/services"

	"github.com/gin-gonic/gin"
)

type ProductController interface {
	FindAll(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewproductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (p *productController) FindAll(context *gin.Context) {
	var products []entity.Product = p.productService.FindAll()
	context.JSON(200, products)
}

func (p *productController) Create(context *gin.Context) {
	var toCreate entity.Product
	err := context.BindJSON(&toCreate)
	if err != nil {
		context.JSON(400, "error creating product")
	} else {
		result := p.productService.Create(toCreate)
		context.JSON(200, result)
	}
}

func (p *productController) Update(context *gin.Context) {
	id := context.Param("id")
	price := context.PostForm("price")
	result := p.productService.Update(id, price)
	context.JSON(200, result)
}
func (p *productController) Delete(context *gin.Context) {
	id := context.Param("id")
	result := p.productService.Delete(id)
	context.JSON(200, result)
}
