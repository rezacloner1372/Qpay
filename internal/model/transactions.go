package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transactions struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	// DeletedAt       *time.Time      `gorm:"index"`
	GatewayID       uint            `gorm:"not null"`
	Gateway         PaymentGateways `gorm:"foreignKey:GatewayID"`
	Amount          int             `gorm:"not null"`
	Status          string          `gorm:"not null"`
	TransactionTime time.Time       `gorm:"not null"`
	UserID          uint            `gorm:"not null"`
	User            User            `gorm:"foreignKey:UserID"`
	CallbackURL     string          `gorm:"not null"`
	Description     string          `gorm:""`
	Email           string          `gorm:""`
	Phone           string          `gorm:"not null"`
	Authority       string          `gorm:"not null"`
}
