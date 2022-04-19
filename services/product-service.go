package services

import (
	"gin-api/dao"
	"gin-api/entity"
)

type ProductService interface {
	FindAll() []entity.Product
	Create(body entity.Product) string
	Update(id string, body entity.Product) string
	Delete(id string) string
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

func (service *productService) Create(body entity.Product) string {
	return service.productDAO.Create(body)
}

func (service *productService) Update(id string, body entity.Product) string {
	return service.productDAO.Update(id, body)
}

func (service *productService) Delete(id string) string {
	return service.productDAO.Delete(id)
}
