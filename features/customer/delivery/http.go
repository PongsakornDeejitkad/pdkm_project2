package delivery

import (
	"log"
	"net/http"
	"order-management/domain"
	"order-management/entity"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CustomerHandler struct {
	usecase domain.CustomerUsecase
}

func NewHandler(e *echo.Group, u domain.CustomerUsecase) *CustomerHandler {
	h := CustomerHandler{usecase: u}
	e.POST("/customer/signin", h.CreateCustomer)
	e.GET("/customer", h.ListCustomers)
	e.GET("/customer/:id", h.GetCustomer)
	e.DELETE("/customer/:id", h.DeleteCustomer)
	e.PUT("/customer/:id", h.UpdateCustomer)
	e.POST("/customer/login", h.CustomerLogin)
	e.POST("/customer/refreshtoken", h.RefreshRequest)

	return &h
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	customer := entity.Customer{}
	c.Bind(&customer)

	validate := validator.New(validator.WithRequiredStructEnabled())
	validateError := validate.Struct(customer)

	if validateError != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation error",
			"errors":  validateError.Error(),
		})
	}

	err := h.usecase.CreateCustomer(customer)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.NoContent(http.StatusCreated)
}

func (h *CustomerHandler) ListCustomers(c echo.Context) error {

	customers, err := h.usecase.ListCustomers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, customers)
}

func (h *CustomerHandler) GetCustomer(c echo.Context) error {
	customerIdString := c.Param("id")
	customerId, err := strconv.Atoi(customerIdString)
	if err != nil {
		log.Println("Invalid customer ID:", customerId)
		return err
	}

	customer, err := h.usecase.GetCustomer(customerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, customer)
}
func (h *CustomerHandler) DeleteCustomer(c echo.Context) error {
	customerIdString := c.Param("id")
	customerId, err := strconv.Atoi(customerIdString)
	if err != nil {
		log.Println("Invalid customer ID:", customerId)
		return err
	}
	usecaseError := h.usecase.DeleteCustomer(customerId)
	if usecaseError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": usecaseError,
		})
	}
	return c.NoContent(http.StatusOK)

}

func (h *CustomerHandler) UpdateCustomer(c echo.Context) error {
	customer := entity.Customer{}
	customerIdString := c.Param("id")
	customerId, err := strconv.Atoi(customerIdString)
	if err != nil {
		log.Println("Invalid customer ID:", customerId)
		return err
	}
	c.Bind(&customer)

	validate := validator.New(validator.WithRequiredStructEnabled())
	validateError := validate.Struct(customer)

	if validateError != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation error",
			"errors":  validateError.Error(),
		})
	}

	usecaseError := h.usecase.UpdateCustomer(customerId, customer)
	if usecaseError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": usecaseError,
		})
	}
	return c.NoContent(http.StatusOK)

}

func (h *CustomerHandler) CustomerLogin(c echo.Context) error {
	customerReq := entity.CustomerLoginRequest{}
	c.Bind(&customerReq)

	validate := validator.New(validator.WithRequiredStructEnabled())
	validateError := validate.Struct(customerReq)

	if validateError != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation error",
			"errors":  validateError.Error(),
		})
	}

	customerRes, err := h.usecase.CustomerLogin(customerReq)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusForbidden, map[string]interface{}{
				"message": "email not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, customerRes)
}

func (h *CustomerHandler) RefreshRequest(c echo.Context) error {
	RefreshRequestToken := entity.RefreshTokenRequest{}
	c.Bind(&RefreshRequestToken)

	customerRes, _ := h.usecase.RefreshRequest(RefreshRequestToken)

	return c.JSON(http.StatusOK, customerRes)
}
