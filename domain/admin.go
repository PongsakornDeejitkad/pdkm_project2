package domain

import "order-management/entity"

type AdminUsecase interface {
	CreateAdmin(admin entity.Admin) error
	ListAdmins() ([]entity.Admin, error)
}

type AdminRepository interface {
	CreateAdmin(admin entity.Admin) error
	ListAdmins() ([]entity.Admin, error)
}
