package services

import (
	"gin-api/dao"
	"gin-api/entity"
)

type ProductService interface {
	FindAll() []entity.Product
}

type productService struct {
	productDAO dao.ProductDAO
}

func NewProductService(productDAO dao.ProductDAO) ProductService {
	return &productService{
		productDAO: productDAO,
	}
}

func (service *productService) FindAll() []entity.Product {
	return service.productDAO.FindAll()
}
