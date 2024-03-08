package entity

type Auth struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Email    string `json:"email" gorm:"not null; unique; varchar(100)" validate:"required,email"`
	Password string `json:"password" gorm:"not null; type:text;size:200" validate:"required"`
}
