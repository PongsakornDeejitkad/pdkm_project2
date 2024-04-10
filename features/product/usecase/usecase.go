package usecase

import (
	"order-management/domain"
	"order-management/entity"
)

type productUsecase struct {
	productRepo domain.ProductRepository
}

func NewUsecase(productRepo domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
	}
}

func (u *productUsecase) CreateProduct(product entity.Product) error {
	return u.productRepo.CreateProduct(product)
}
