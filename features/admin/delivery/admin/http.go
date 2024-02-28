package delivery

import (
	"net/http"
	"order-management/domain"
	"order-management/entity"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	usecase domain.AdminUsecase
}

func NewAdminHandler(e *echo.Group, u domain.AdminUsecase) *AdminHandler {
	h := AdminHandler{usecase: u}

	e.GET("/test", h.TestAdmin)
	e.POST("", h.CreateAdmin)

	return &h
}

func (h *AdminHandler) CreateAdmin(c echo.Context) error {
	admin := entity.Admin{}
	c.Bind(&admin)

	err := h.usecase.CreateAdmin(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.NoContent(http.StatusCreated)
}

func (h *AdminHandler) TestAdmin(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"status": "admin test success"})
}
