package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID           int            `json:"id" gorm:"primary_key"`
	Name         string         `json:"name" gorm:"not null; varchar(100)"`
	Descriptions string         `json:"description" gorm:"text" `
	SKU          int            `json:"sku" gorm:"integer"`
	Price        float64        `json:"price" gorm:"float"`
	Cost         float64        `json:"cost" gorm:"float"`
	CategoryID   int            `json:"category_id" gorm:"integer"`
	Quantity     int            `json:"quantity" gorm:"integer"`
	CreatedAt    time.Time      `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"default: NOW()"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Category ProductsCategory `json:"-" gorm:"foreignKey:CategoryID"`
}

type ProductsCategory struct {
	ID          int            `json:"id" gorm:"primary_key"`
	Name        string         `json:"name" gorm:"varchar(100)"`
	Description string         `json:"description" gorm:"type:text;size:255"`
	CreatedAt   time.Time      `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"default: NOW()"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Products []Product `json:"-" gorm:"foreignKey:CategoryID"`
}
