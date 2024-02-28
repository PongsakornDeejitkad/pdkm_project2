package delivery

import (
	"order-management/domain"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase domain.AdminUsecase
}

func NewHandler(e *echo.Group, u domain.AdminUsecase) *Handler {
	h := Handler{usecase: u}

	// e.POST("/admin", h.CreateAdmin)

	return &h
}
