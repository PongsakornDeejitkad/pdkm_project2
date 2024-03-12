package usecase

import (
	"errors"
	"order-management/domain"
	"order-management/entity"
	"os"
	"time"

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

	accessTokenClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	secretKey := []byte(os.Getenv("TOKEN_SECRET"))
	tokenString, tokenErr := token.SignedString(secretKey)
	if tokenErr != nil {
		return customerRes, tokenErr
	}

	customerRes.AccessToken = tokenString

	refreshClaims := jwt.StandardClaims{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	refreshTokenString, refreshTokenErr := refreshToken.SignedString(secretKey)
	if refreshTokenErr != nil {
		return customerRes, refreshTokenErr
	}

	customerRes.RefreshToken = refreshTokenString

	return customerRes, nil
}
