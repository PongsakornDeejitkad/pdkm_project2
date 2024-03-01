package repository

import (
	"log"
	"order-management/domain"
	"order-management/entity"
	"strconv"

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
func (r *adminRepository) ListAdmins() ([]entity.Admin, error) {
	admins := []entity.Admin{}
	if err := r.db.Find(&admins).Error; err != nil {
		log.Println("ListAdmins error: ", err)
		return nil, err
	}
	return admins, nil
}

func (r *adminRepository) GetAdmin(id string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	adminId, _ := strconv.Atoi(id)
	if err := r.db.First(admin, adminId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Admin not found")
			return nil, err
		}
		log.Println("Error retrieving admin:", err)
		return nil, err
	}
	return admin, nil
}

// func (r *adminRepository) UpdateAdmin(admin entity.Admin)
