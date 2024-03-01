package domain

import "order-management/entity"

type AdminUsecase interface {
	CreateAdmin(admin entity.Admin) error
	ListAdmins() ([]entity.Admin, error)
	GetAdmin(id string) (*entity.Admin, error)
	UpdateAdmin(id string, admin entity.Admin) error
	DeleteAdmin(id string) error
}

type AdminRepository interface {
	CreateAdmin(admin entity.Admin) error
	ListAdmins() ([]entity.Admin, error)
	GetAdmin(id string) (*entity.Admin, error)
	UpdateAdmin(id string, admin entity.Admin) error
	DeleteAdmin(id string) error
}
