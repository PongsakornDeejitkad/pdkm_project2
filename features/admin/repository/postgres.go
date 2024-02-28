package repository

import (
	"log"
	"order-management/domain"
	"order-management/entity"

	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) domain.AdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r *adminRepository) CreateAdmin(admin entity.Admin) error {
	if err := r.db.Create(&admin).Error; err != nil {
		log.Println("CreateAdmin error: ", err)
		return err
	}

	return nil
}
