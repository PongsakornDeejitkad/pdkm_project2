package middleware

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AdminAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			errMessage := "Authentication failed."

			ignoreRoutes := []string{"/admin/v1/login"}
			for _, v := range ignoreRoutes {
				if c.Request().RequestURI == v {
					return next(c)
				}
			}

			token := c.Request().Header.Get("Authorization")
			if token == "admin" {
				return next(c)
			}

			if token == "" {
				return c.JSON(401, map[string]interface{}{
					"message": errMessage,
				})
			}

			// Token Validator

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

			ignoreRoutes := []string{"/v1/customer/login", "/v1/customer/signin"}
			for _, v := range ignoreRoutes {
				if c.Request().RequestURI == v {
					return next(c)
				}
			}

			accessToken := c.Request().Header.Get("Authorization")
			if accessToken == "admin" {
				return next(c)
			}

			if accessToken == "" {
				return c.JSON(401, map[string]interface{}{
					"message": errMessage,
				})
			}

			claims := jwt.MapClaims{}
			secretKey := os.Getenv("key.secretKey")
			token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
				log.Println("token", token)
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				return c.JSON(401, map[string]interface{}{
					"message": errMessage,
				})
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				c.Set("user_id", claims["user_id"])
				c.Set("username", claims["username"])
				return next(c)
			}

			c.Set("token", token)

			return next(c)
		}
	}
}
