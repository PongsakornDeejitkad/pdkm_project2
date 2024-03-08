package usecase

import (
	"errors"
	"order-management/domain"
	"order-management/entity"
)

type customerUsecase struct {
	customerRepo domain.CustomerRepository
}

func NewCustomerUsecase(customerRepo domain.CustomerRepository) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (u *customerUsecase) CreateCustomer(customer entity.Customer) error {
	// TODO: Password hashing here!

	return u.customerRepo.CreateCustomer(customer)
}

func (u *customerUsecase) ListCustomers() ([]entity.Customer, error) {
	return u.customerRepo.ListCustomers()
}

func (u *customerUsecase) GetCustomer(id int) (*entity.Customer, error) {
	return u.customerRepo.GetCustomer(id)
}
func (u *customerUsecase) DeleteCustomer(id int) error {
	return u.customerRepo.DeleteCustomer(id)
}

func (u *customerUsecase) UpdateCustomer(id int, customer entity.Customer) error {
	return u.customerRepo.UpdateCustomer(id, customer)
}

func (u *customerUsecase) CustomerLogin(customerReq entity.CustomerLoginRequest) (entity.CustomerLoginResponse, error) {
	customerRes := entity.CustomerLoginResponse{}

	customer, err := u.customerRepo.GetCustomerByEmail(customerReq.Email)
	if err != nil {
		return customerRes, err
	}

	if customer.ID == 0 {
		return customerRes, errors.New("email or password is incorrect")
	}

	// TODO: Password validation here!

	// TODO: Claim JWT here!

	// MOCK
	customerRes.AccessToken = "zzyy123456"

	return customerRes, nil
}
