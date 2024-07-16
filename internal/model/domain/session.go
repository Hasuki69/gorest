package domain

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	Token     string    `json:"token" gorm:"unique"`
	ExpiredAt time.Time `json:"expired_at"`
}
