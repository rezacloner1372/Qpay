package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tariffs struct {
	gorm.Model
	ID           uint      `gorm:"primaryKey"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time
	Name         string `gorm:"not null"`
	Description  string `gorm:"not null"`
	Price        int    `gorm:"not null"`
	Currency     string `gorm:"not null"`
	ValidityDays int    `gorm:"not null"`
	IsDefault    int    `gorm:"default:0"`
}
