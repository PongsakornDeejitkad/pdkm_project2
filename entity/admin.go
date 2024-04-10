package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Admin struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Username  string         `json:"user_name" gorm:"not null;varchar(50)" validate:"required"`
	Password  string         `json:"password" gorm:"not null;type:text;size:200" validate:"required,min=8"`
	FirstName string         `json:"first_name" gorm:"varchar;not null" validate:"required"`
	LastName  string         `json:"last_name" gorm:"varchar;not null" validate:"required"`
	Email     string         `json:"email" gorm:"varchar(100);not null;unique" validate:"required,email"`
	TypeID    int            `json:"type_id" gorm:"integer"`
	LastLogin time.Time      `json:"last_login" gorm:"default: NOW()" `
	CreatedAt time.Time      `json:"created_at" gorm:"default: NOW()" `
	UpdatedAt time.Time      `json:"updated_at" gorm:"default: NOW()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Type AdminType `json:"-" gorm:"foreignKey:TypeID"`
}

type AdminType struct {
	ID         int       `json:"id" gorm:"primary_key"`
	AdminType  string    `json:"admin_type" gorm:"varchar" `
	Permission string    `json:"permission" gorm:"varchar"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:NOW()"`

	Admins []Admin `json:"-" gorm:"foreignKey:TypeID"`
}

type AdminLoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	jwt.StandardClaims
}
