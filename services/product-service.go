package services

import (
	"gin-api/entity"
	"gin-api/repo"
)

type ProductService interface {
	FindAll() []entity.Product
}

type productService struct {
	productRepo repo.ProductRepo
}

func NewProductService(productRepo repo.ProductRepo) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (service *productService) FindAll() []entity.Product {
	return service.productRepo.FindAll()
}
