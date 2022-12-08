package database

import (
	"sample/git/models"

	"github.com/jinzhu/gorm"
)

type DbClient struct {
	Db *gorm.DB
}

type DbProducts interface {
	GetProducts() ([]models.Product, error)
}

func (s DbClient) GetProducts() ([]models.Product, error) {
	var products []models.Product
	s.Db.Find(&products)
	return products, nil
}
