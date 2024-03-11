package domain

import "order-management/entity"

type AdminUsecase interface {
	CreateAdmin(admin entity.Admin) error
	CreateAdminType(adminType entity.AdminType) error
	ListAdmins() ([]entity.Admin, error)
	ListAdminTypes() ([]entity.AdminType, error)
	GetAdmin(id int) (*entity.Admin, error)
	UpdateAdmin(id int, admin entity.Admin) error
	DeleteAdmin(id int) error

	// AdminLogin(entity.AdminLoginRequest) (entity.AdminLoginResponse, error)
}

type AdminRepository interface {
	CreateAdmin(admin entity.Admin) error
	CreateAdminType(adminType entity.AdminType) error
	ListAdmins() ([]entity.Admin, error)
	ListAdminTypes() ([]entity.AdminType, error)
	GetAdmin(id int) (*entity.Admin, error)
	UpdateAdmin(id int, admin entity.Admin) error
	DeleteAdmin(id int) error

	// GetAdminByEmail(email string) (*entity.Admin, error)
}
