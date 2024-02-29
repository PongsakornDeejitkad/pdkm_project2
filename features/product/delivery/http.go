package delivery

import (
	"net/http"
	"order-management/domain"
	"order-management/entity"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	usecase domain.ProductUsecase
}

func NewProductHandler(e *echo.Group, u domain.ProductUsecase) *ProductHandler {
	h := ProductHandler{usecase: u}

	// e.GET("/", h.TestProduct)
	e.POST("", h.CreateProduct)

	return &h
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := entity.Product{}
	c.Bind(&product)

	err := h.usecase.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.NoContent(http.StatusCreated)
}
