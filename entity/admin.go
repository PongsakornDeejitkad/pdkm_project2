package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Username  string         `json:"user_name" gorm:"not null;varchar(50)"`
	Password  string         `json:"password" gorm:"not null;type:text;size:200"`
	FirstName string         `json:"first_name" gorm:"varchar;not null"`
	LastName  string         `json:"last_name" gorm:"varchar;not null"`
	Email     string         `json:"email" gorm:"varchar(100);not null;unique"`
	TypeID    int            `json:"type_id" gorm:"integer"`
	LastLogin time.Time      `json:"last_login" gorm:"default: NOW()"`
	CreatedAt time.Time      `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default: NOW()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Type AdminType `json:"-" gorm:"foreignKey:TypeID"`
}

type AdminType struct {
	ID         int       `json:"id" gorm:"primary_key"`
	AdminType  string    `json:"admin_type" gorm:"varchar"`
	Permission string    `json:"permission" gorm:"varchar"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:NOW()"`

	// Admins []Admin `json:"-" gorm:"foreignKey:TypeID"`
}
