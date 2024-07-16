package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Username string  `json:"username" gorm:"unique"`
	Password string  `json:"password"`
	Session  Session `json:"session"`
}
