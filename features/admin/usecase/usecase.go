package usecase

import (
	"order-management/domain"
	"order-management/entity"
)

type adminUsecase struct {
	adminRepo domain.AdminRepository
}

func NewAdminUsecase(adminRepo domain.AdminRepository) domain.AdminUsecase {
	return &adminUsecase{
		adminRepo: adminRepo,
	}
}

func (u *adminUsecase) CreateAdmin(admin entity.Admin) error {
	return u.adminRepo.CreateAdmin(admin)
}

func (u *adminUsecase) ListAdmins() ([]entity.Admin, error) {
	return u.adminRepo.ListAdmins()
}

func (u *adminUsecase) GetAdmin(id string) (*entity.Admin, error) {
	return u.adminRepo.GetAdmin(id)
}
func (u *adminUsecase) UpdateAdmin(id string, admin entity.Admin) error {
	return u.adminRepo.UpdateAdmin(id, admin)
}

func (u *adminUsecase) DeleteAdmin(id string) error {
	return u.adminRepo.DeleteAdmin(id)
}

func (u *adminUsecase) CreateAdminType(adminType entity.AdminType) error {
	return u.adminRepo.CreateAdminType(adminType)
}

func (u *adminUsecase) ListAdminTypes() ([]entity.AdminType, error) {
	return u.adminRepo.ListAdminTypes()
}
