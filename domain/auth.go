package domain

import "order-management/entity"

type AuthUsecase interface {
	FindOneUserByEmail(auth entity.Auth) string
}

type AuthRepository interface {
	FindOneUserByEmail(auth entity.Auth) error
}
