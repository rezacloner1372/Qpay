package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type PaymentGateways struct {
	gorm.Model
	Name              string    `gorm:"not null"`
	UserId            string    `gorm:"not null"`
	PersonalizedUrl   string    `gorm:"not null"`
	Tariff            int       `grom:"not null"`
	Status            int       `grom:"not null"`
	RegisterationDate time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
	User              User      `gorm:"references:id"`
	Tariffs           Tariffs   `gorm:"references:id"`
}
