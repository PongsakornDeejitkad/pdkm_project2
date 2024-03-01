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
	adminId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid admin ID")
		return nil, err
	}
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

func (r *adminRepository) UpdateAdmin(id string, admin entity.Admin) error {
	adminId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Invalid admin ID")
		return err
	}

	existingAdmin := &entity.Admin{}
	if err := r.db.First(&existingAdmin, adminId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Admin not found")
			return err
		}
		log.Println("Error retrieving admin:", err)
		return err
	}
	if err := r.db.Model(&entity.Admin{}).Where("id = ?", adminId).Updates(&admin).Error; err != nil {
		log.Println("Error updating admin:", err)
		return err
	}

	return nil
}

func (r *adminRepository) DeleteAdmin(id string) error {
	adminId, err := strconv.Atoi(id)

	if err != nil {
		log.Println("Invalid admin ID:", id)
		return err
	}
	existingAdmin := &entity.Admin{}
	if err := r.db.First(&existingAdmin, adminId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Admin not found")
			return err
		}
		log.Println("Error retrieving admin:", err)
		return err
	}

	if err := r.db.Model(&entity.Admin{}).Where("id = ?", adminId).Delete(&existingAdmin).Error; err != nil {
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
