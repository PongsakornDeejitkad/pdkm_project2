package repositories

import (
	"log"
	"order-management/domain"
	"order-management/entity"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &authRepo{
		db: db,
	}
}

func (r *authRepo) FindOneUserByEmail(auth entity.Auth) error {
	existingUser := &entity.Customer{}

	if err := r.db.Model(&entity.Customer{}).Where("email = ? AND password = ?", auth.Email, auth.Password).First(&existingUser).Error; err != nil {
		log.Println("Error Finding user", err)
		return err
	}

	return nil

}
