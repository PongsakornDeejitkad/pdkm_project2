package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Customer struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" gorm:"not null; varchar(100)" validate:"required"`
	Email     string         `json:"email" gorm:"not null; unique; varchar(100)" validate:"required,email"`
	Username  string         `json:"user_name" gorm:"not null; unique; varchar(50)" validate:"required"`
	Password  string         `json:"password" gorm:"not null; type:text;size:200" validate:"required"`
	CreatedAt time.Time      `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default: NOW()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CustomerLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CustomerLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	jwt.StandardClaims
}
