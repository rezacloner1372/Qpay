package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type PaymentGateways struct {
	gorm.Model
	Id                int
	Name              string    `gorm:"not null"`
	UserId            string    `gorm:"not null"`
	PersonalizedUrl   string    `gorm:"not null"`
	Tariff            int       `grom:"not null"`
	Status            int       `grom:"not null"`
	RegisterationDate time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
}
