package usecase

import (
	"errors"
	"order-management/domain"
	"order-management/entity"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	customer.Password = string(hashedPassword)
	return u.customerRepo.UpdateCustomer(id, customer)
}

func (u *customerUsecase) CustomerLogin(customerReq entity.CustomerLoginRequest) (entity.CustomerLoginResponse, error) {
	customerRes := entity.CustomerLoginResponse{}

	customer, err := u.customerRepo.GetCustomerByEmail(customerReq.Email)
	if err != nil {
		return customerRes, err
	}

	if customer.ID == 0 {
		return customerRes, errors.New("email is incorrect")
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(customerReq.Password))
	if passwordErr != nil {
		return customerRes, errors.New("password is incorrect")
	}

	// Generate access token
	claims := entity.CustomerClaims{
		Id:       customer.ID,
		Username: customer.Username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	secretKey := []byte(os.Getenv("key.secretKey"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, tokenErr := token.SignedString(secretKey)
	if tokenErr != nil {
		return customerRes, tokenErr
	}

	//  Generate refresh token
	refreshToken := entity.RefreshRequest{
		RefreshToken: uuid.New().String(),
		CustomerID:   customer.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	customerRes.AccessToken = tokenString
	customerRes.RefreshToken = refreshToken.RefreshToken

	return customerRes, nil
}
