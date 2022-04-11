package config

import (
	"gin-api/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDB() *gorm.DB {
	dsn := "host=localhost user=arpit password=leapfrog123 dbname=crudapi port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{})
	return db
}
