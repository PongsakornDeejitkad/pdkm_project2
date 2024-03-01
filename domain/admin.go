package domain

import "order-management/entity"

type AdminUsecase interface {
	CreateAdmin(admin entity.Admin) error
	ListAdmins() ([]entity.Admin, error)
	GetAdmin(id string) (*entity.Admin, error)
}

type AdminRepository interface {
	CreateAdmin(admin entity.Admin) error
	ListAdmins() ([]entity.Admin, error)
	GetAdmin(id string) (*entity.Admin, error)
}
