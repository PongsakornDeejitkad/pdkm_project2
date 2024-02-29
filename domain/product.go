package domain

import "order-management/entity"

type ProductUsecase interface {
	CreateProduct(product entity.Product) error
}

type ProductRepository interface {
	CreateProduct(product entity.Product) error
}
