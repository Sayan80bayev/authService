package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Password string `json:"password"`
	Role     Role   `json:"role" gorm:"not null;default:'USER'"`
	Active   bool   `json:"active" gorm:"default:true"`
}
