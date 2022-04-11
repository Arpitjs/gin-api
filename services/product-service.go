package services

import (
	"gin-api/dao"
	"gin-api/entity"
)

type ProductService interface {
	FindAll() []entity.Product
	Create(body entity.Product) string
	Update(id string, price string) string
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
	res := service.productDAO.Create(body)
	return res
}
func (service *productService) Update(id string, price string) string {
	res := service.productDAO.Update("price", price)
	return res
}
func (service *productService) Delete(id string) string {
	res := service.productDAO.Delete(id)
	return res
}
