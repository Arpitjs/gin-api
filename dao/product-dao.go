package dao

import (
	"gin-api/entity"

	"gorm.io/gorm"
)

type ProductDAO interface {
	FindAll() []entity.Product
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
