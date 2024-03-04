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

type AdminHandler struct {
	usecase domain.AdminUsecase
}

func NewAdminHandler(e *echo.Group, u domain.AdminUsecase) *AdminHandler {
	h := AdminHandler{usecase: u}

	e.GET("", h.ListAdmins)
	e.GET("/types", h.ListAdminTypes)
	e.GET("/:id", h.GetAdmin)
	e.POST("", h.CreateAdmin)
	e.POST("/types", h.CreateAdminType)
	e.PUT("/:id", h.UpdateAdmin)
	e.DELETE("/:id", h.DeleteAdmin)

	return &h
}

func (h *AdminHandler) CreateAdmin(c echo.Context) error {
	admin := entity.Admin{}
	c.Bind(&admin)

	validate := validator.New(validator.WithRequiredStructEnabled())
	validateError := validate.Struct(admin)

	if validateError != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation error",
			"errors":  validateError.Error(),
		})
	}

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
	adminId, err := strconv.Atoi(adminIdString)
	if err != nil {
		log.Println("Invalid admin ID")
		return err
	}
	admin, err_usecase := h.usecase.GetAdmin(adminId)
	if err_usecase != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err_usecase,
		})
	}
	return c.JSON(http.StatusOK, admin)
}

func (h *AdminHandler) UpdateAdmin(c echo.Context) error {
	admin := entity.Admin{}
	adminIdString := c.Param("id")
	adminId, err := strconv.Atoi(adminIdString)
	if err != nil {
		log.Println("Invalid admin ID")
		return err
	}
	c.Bind(&admin)

	validate := validator.New(validator.WithRequiredStructEnabled())
	validateError := validate.Struct(admin)

	if validateError != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation error",
			"errors":  validateError.Error(),
		})
	}

	err_usecase := h.usecase.UpdateAdmin(adminId, admin)
	if err_usecase != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err_usecase,
		})
	}
	return c.NoContent(http.StatusOK)
}

func (h *AdminHandler) DeleteAdmin(c echo.Context) error {
	adminIdString := c.Param("id")
	adminId, err := strconv.Atoi(adminIdString)
	if err != nil {
		log.Println("Invalid admin ID:", adminId)
		return err
	}

	usecaseError := h.usecase.DeleteAdmin(adminId)
	if usecaseError != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": usecaseError,
		})
	}
	return c.NoContent(http.StatusOK)
}

func (h *AdminHandler) CreateAdminType(c echo.Context) error {
	adminType := entity.AdminType{}
	c.Bind(&adminType)

	validate := validator.New(validator.WithRequiredStructEnabled())
	validateError := validate.Struct(adminType)

	if validateError != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Validation error",
			"errors":  validateError.Error(),
		})
	}

	err := h.usecase.CreateAdminType(adminType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.NoContent(http.StatusCreated)

}

func (h *AdminHandler) ListAdminTypes(c echo.Context) error {
	adminTypes, err := h.usecase.ListAdminTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, adminTypes)
}
