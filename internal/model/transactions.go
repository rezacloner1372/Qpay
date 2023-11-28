package model

import (
	"time"
)

type Transactions struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt       time.Time `gorm:"index"`
	UpdatedAt       time.Time
	DeletedAt       *time.Time      `gorm:"index"`
	GatewayID       uint            `gorm:"not null"`
	Gateway         PaymentGateways `gorm:"foreignKey:GatewayID"`
	Amount          int             `gorm:"not null"`
	Status          string          `gorm:"not null"`
	TransactionTime time.Time       `gorm:"not null"`
	UserID          uint            `gorm:"not null"`
	User            User            `gorm:"foreignKey:UserID"`
}
