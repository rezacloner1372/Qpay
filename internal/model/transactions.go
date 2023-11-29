package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Transactions struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time
	GatewayID       uint            `gorm:"not null"`
	Gateway         PaymentGateways `gorm:"foreignKey:GatewayID"`
	Amount          int             `gorm:"not null"`
	Status          string          `gorm:"not null"`
	TransactionTime time.Time       `gorm:"default:0000-00-00 00:00:00"`
	UserID          uint            `gorm:"not null"`
	User            User            `gorm:"foreignKey:UserID"`
	CallbackURL     string          `gorm:"not null"`
	Description     string          `gorm:""`
	Email           string          `gorm:""`
	Phone           string          `gorm:"not null"`
	Authority       string          `gorm:"unique_index;not null"`
	BankAuthority   string          `gorm:"unique_index;not null"`
	BankRefID       string          `gorm:""`
}
