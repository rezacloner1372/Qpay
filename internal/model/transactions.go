package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Transactions struct {
	gorm.Model
	GatewayId       uint      `gorm:"column:gateway_id"`
	Amount          string    `gorm:"not null"`
	Status          string    `gorm:"not null"`
	TransactionTime time.Time `gorm:"not null"`
	UserId          uint      `gorm:"column:user_id"`
}
