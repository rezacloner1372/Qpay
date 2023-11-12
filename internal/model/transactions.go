package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Transactions struct {
	gorm.Model
	GatewayId       int       `gorm:"not null"`
	Amount          string    `gorm:"not null"`
	Status          string    `gorm:"not null"`
	TransactionTime time.Time `gorm:"not null"`
	UserId          int       `gorm:"not null"`
}
