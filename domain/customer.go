package domain

import "order-management/entity"

type CustomerUsecase interface {
	CreateCustomer(customer entity.Customer) error
	ListCustomers() ([]entity.Customer, error)
	GetCustomer(id int) (*entity.Customer, error)
	DeleteCustomer(id int) error
	UpdateCustomer(id int, customer entity.Customer) error

	CustomerLogin(entity.CustomerLoginRequest) (entity.CustomerLoginResponse, error)
	RefreshRequest(RefreshRequest entity.RefreshTokenRequest) (entity.CustomerLoginResponse, error)
}

type CustomerRepository interface {
	CreateCustomer(customer entity.Customer) error
	ListCustomers() ([]entity.Customer, error)
	GetCustomer(id int) (*entity.Customer, error)
	DeleteCustomer(id int) error
	UpdateCustomer(id int, customer entity.Customer) error

	GetCustomerByEmail(email string) (entity.Customer, error)
}
