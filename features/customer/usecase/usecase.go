package usecase

import (
	"errors"
	"log"
	"order-management/domain"
	"order-management/entity"
	"os"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type customerUsecase struct {
	customerRepo domain.CustomerRepository
}

func NewUsecase(customerRepo domain.CustomerRepository) domain.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (u *customerUsecase) CreateCustomer(customer entity.Customer) error {
	// TODO: Password hashing here!
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	customer.Password = string(hashedPassword)

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

	log.Println(customer.Password)
	log.Println(customerReq.Password)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customerRes)
	secretKey := []byte(os.Getenv("TOKEN_SECRET"))
	tokenString, tokenErr := token.SignedString(secretKey)
	if tokenErr != nil {
		return customerRes, tokenErr
	}
	log.Println(tokenString)

	customerRes.AccessToken = tokenString

	return customerRes, nil
}

// passwordErr := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(customerReq.Password))
// if passwordErr != nil {
// 	return customerRes, errors.New("email or password is incorrect")
// }

