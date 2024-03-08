package delivery

import (
	"net/http"
	"order-management/domain"
	"order-management/entity"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(e *echo.Echo, u domain.AuthUsecase) {
	h := &AuthHandler{
		usecase: u,
	}
	e.GET("/auth", h.GetToken)
}

func (h *AuthHandler) GetToken(c echo.Context) error {
	auth := entity.Auth{}

	if err := c.Bind(auth); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	token := h.usecase.FindOneUserByEmail(auth)

	return c.String(http.StatusOK, token)

}
