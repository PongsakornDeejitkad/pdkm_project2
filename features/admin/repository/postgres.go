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
func (r *adminRepository) ListAdmins() ([]entity.Admin, error) {
	admins := []entity.Admin{}
	if err := r.db.Find(&admins).Error; err != nil {
		log.Println("ListAdmins error: ", err)
		return nil, err
	}
	return admins, nil
}

func (r *adminRepository) GetAdmin(id int) (*entity.Admin, error) {
	admin := &entity.Admin{}
	if err := r.db.First(admin, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Admin not found")
			return nil, err
		}
		log.Println("Error retrieving admin:", err)
		return nil, err
	}
	return admin, nil
}

func (r *adminRepository) UpdateAdmin(id int, admin entity.Admin) error {
	existingAdmin := &entity.Admin{}
	if err := r.db.First(&existingAdmin, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Admin not found")
			return err
		}
		log.Println("Error retrieving admin:", err)
		return err
	}
	if err := r.db.Model(&entity.Admin{}).Where("id = ?", id).Updates(&admin).Error; err != nil {
		log.Println("Error updating admin:", err)
		return err
	}

	return nil
}

func (r *adminRepository) DeleteAdmin(id int) error {
	existingAdmin := &entity.Admin{}
	if err := r.db.First(&existingAdmin, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Admin not found")
			return err
		}
		log.Println("Error retrieving admin:", err)
		return err
	}

	if err := r.db.Model(&entity.Admin{}).Where("id = ?", id).Delete(&existingAdmin).Error; err != nil {
		log.Println("Error delete admin:", err)
		return err
	}
	return nil

}

func (r *adminRepository) CreateAdminType(adminType entity.AdminType) error {
	if err := r.db.Create(&adminType).Error; err != nil {
		log.Println("Create AdminType error: ", err)
		return err
	}

	return nil

}

func (r *adminRepository) ListAdminTypes() ([]entity.AdminType, error) {
	adminTypes := []entity.AdminType{}
	if err := r.db.Find(&adminTypes).Error; err != nil {
		log.Println("ListAdminTypes error: ", err)
		return nil, err
	}
	return adminTypes, nil
}
