package repository

import (
	"log"
	"order-management/domain"
	"order-management/entity"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) CreateProduct(product entity.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		log.Println("CreteProduct error :", err)
		return err
	}
	return nil
}
