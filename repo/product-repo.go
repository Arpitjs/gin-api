package repo

import (
	"gin-api/entity"

	"gorm.io/gorm"
)

type ProductRepo interface {
	FindAll() []entity.Product
}

type productConnection struct {
	connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) ProductRepo {
	return &productConnection{
		connection: dbConn,
	}
}

func (db *productConnection) FindAll() []entity.Product {
	var products []entity.Product
	db.connection.Find(&products)
	return products
}
