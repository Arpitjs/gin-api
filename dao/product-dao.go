package dao

import (
	"gin-api/entity"

	"gorm.io/gorm"
)

type ProductDAO interface {
	FindAll() []entity.Product
	Create(body entity.Product) string
	Update(id string, body entity.Product) string
	Delete(id string) string
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductDAO(dbConn *gorm.DB) ProductDAO {
	return &productConnection{
		connection: dbConn,
	}
}

func (db *productConnection) FindAll() []entity.Product {
	var products []entity.Product
	db.connection.Find(&products)
	return products
}

func (db *productConnection) Create(body entity.Product) string {
	db.connection.Create(body)
	return "product created!"
}
func (db *productConnection) Update(id string, body entity.Product) string {
	db.connection.Where("id", id).Updates(body)
	return "product updated!"
}
func (db *productConnection) Delete(id string) string {
	var product entity.Product
	db.connection.Where("id", id).Delete(&product)
	return "product deleted!"
}
