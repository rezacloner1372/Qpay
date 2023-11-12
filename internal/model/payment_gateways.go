package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PaymentGateways struct {
	gorm.Model
	Name              string    `gorm:"not null"`
	UserId            uint      `gorm:"column:user_id"`
	PersonalizedUrl   string    `gorm:"not null"`
	Tariff            int       `gorm:"column:tariff"`
	Status            int       `grom:"not null"`
	RegisterationDate time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
}
