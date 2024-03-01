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

	e.GET("", h.ListAdmins)
	e.POST("", h.CreateAdmin)
	e.GET("/:id", h.GetAdmin)

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

func (h *AdminHandler) ListAdmins(c echo.Context) error {
	admins, err := h.usecase.ListAdmins()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, admins)
}

func (h *AdminHandler) GetAdmin(c echo.Context) error {
	adminIdString := c.Param("id")
	admin, err := h.usecase.GetAdmin(adminIdString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}
	return c.JSON(http.StatusOK, admin)
}
