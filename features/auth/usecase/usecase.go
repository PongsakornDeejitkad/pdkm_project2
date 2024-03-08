package usecases

import (
	"order-management/domain"
	"order-management/entity"
)

type authUsecase struct {
	AuthRepo domain.AuthRepository
}

func NewAuthUsecase(authRepo domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{
		AuthRepo: authRepo,
	}
}

func (u *authUsecase) FindOneUserByEmail(auth entity.Auth) string {

	err := u.AuthRepo.FindOneUserByEmail(auth)
	if err != nil {
		// genToken
	}

	return "genToken"
}
