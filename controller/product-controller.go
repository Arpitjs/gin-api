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

type Map map[string]interface{}

func NewproductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

// GetProducts godoc
// @Summary List existing products
// @Description Get all the existing products
// @Tags code, price
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Product
// @Failure 400 {string} "not found"
// @Router /products [get]
func (p *productController) FindAll(context *gin.Context) {
	var products []entity.Product = p.productService.FindAll()
	context.JSON(200, products)
}

// CreateProduct godoc
// @Summary Create new products
// @Description Create a new product
// @Tags code, price
// @Accept  json
// @Produce  json
// @Param product body entity.Product true "Create product"
// @Success 200 {string} "created product successfully..."
// @Failure 400 {string} "error creating product"
// @Router /products [post]
func (p *productController) Create(context *gin.Context) {
	var toCreate entity.Product
	err := context.BindJSON(&toCreate)
	if err != nil {
		context.JSON(400, "error creating product")
		return
	} else {
		result := p.productService.Create(toCreate)
		m := Map{"info": result, "data": toCreate}
		context.JSON(200, m)
	}
}

func (p *productController) Update(context *gin.Context) {
	id := context.Param("id")
	var toUpdate entity.Product
	err := context.BindJSON(&toUpdate)
	if err != nil {
		context.JSON(400, "error creating product")
	} else {
		result := p.productService.Update(id, toUpdate)
		m := Map{"info": result, "data": toUpdate}
		context.JSON(200, m)
	}
}

func (p *productController) Delete(context *gin.Context) {
	id := context.Param("id")
	result := p.productService.Delete(id)
	context.JSON(200, result)
}
