package delivery

import (
	"log"
	"net/http"
	"order-management/domain"
	"order-management/entity"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	usecase domain.CustomerUsecase
}

func NewCustomerHandler(e *echo.Group, u domain.CustomerUsecase) *CustomerHandler {
	h := CustomerHandler{usecase: u}
	e.POST("", h.CreateCustomer)
	e.GET("", h.ListCustomers)
	e.GET("/:id", h.GetCustomer)
	e.DELETE("/:id", h.DeleteCustomer)
	e.PUT("/:id", h.UpdateCustomer)

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

	usecaseError := h.usecase.UpdateCustomer(customerId, customer)
	if usecaseError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": usecaseError,
		})
	}
	return c.NoContent(http.StatusOK)

}
