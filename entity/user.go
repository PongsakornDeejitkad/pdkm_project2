package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" gorm:"not null; varchar(100)"`
	Email     string         `json:"email" gorm:"not null; unique; varchar(100)"`
	Username  string         `json:"username" gorm:"not null; unique; varchar(50)"`
	Password  string         `json:"password" gorm:"not null; type:text;size:200"`
	CreatedAt time.Time      `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default: NOW()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
