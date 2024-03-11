package middleware

import (
	"github.com/labstack/echo/v4"
)

func AdminAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			errMessage := "Authentication failed."

			ignoreRoutes := []string{"/admin/v1/login", "/admin/v1/app-setting/tax-year"}
			for _, v := range ignoreRoutes {
				if c.Request().RequestURI == v {
					return next(c)
				}
			}

			token := c.Request().Header.Get("Authorization")

			if token != "zzyy123456" || token == "" {
				return c.JSON(401, map[string]interface{}{
					"message": errMessage,
				})
			}

			// c.Set("email", claims["email"])
			// c.Set("name", claims["name"])
			// c.Set("surname", claims["surname"])
			// c.Set("permissions", claims["permissions"])
			// c.Set("role_id", claims["role_id"])
			// c.Set("department_id", claims["department_id"])

			return next(c)
		}
	}
}

func CustomerAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			errMessage := "Authentications failed."

			ignoreRoutes := []string{"/v1/customer/login", "/v1/customers", "/v1/customer"}
			for _, v := range ignoreRoutes {
				if c.Path() == v {
					return next(c)
				}
			}

			token := c.Request().Header.Get("Authorization")

			if token != "zzyy123456" || token == "" {
				return c.JSON(401, map[string]interface{}{
					"message": errMessage,
				})
			}

			// c.Set("email", claims["email"])
			// c.Set("name", claims["name"])
			// c.Set("surname", claims["surname"])
			// c.Set("permissions", claims["permissions"])
			// c.Set("role_id", claims["role_id"])
			// c.Set("department_id", claims["department_id"])

			return next(c)
		}
	}
}
