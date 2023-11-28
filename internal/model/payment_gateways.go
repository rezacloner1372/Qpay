package model

import (
	"time"
)

type PaymentGateways struct {
	ID              uint       `gorm:"primaryKey"`
	CreatedAt       time.Time  `gorm:"autoCreateTime"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime"`
	DeletedAt       *time.Time `gorm:"index"`
	Title           string     `gorm:"not null"`
	UserID          uint       `gorm:"not null"`
	User            User       `gorm:"foreignKey:UserID"`
	PersonalizedURL string     `gorm:"not null"`
	IsDefault       int        `gorm:"not null;default:1"`
	TariffID        *uint      // Assuming tariffs are pointers (optional field)
	Tariff          *Tariffs   `gorm:"foreignKey:TariffID"`
}
