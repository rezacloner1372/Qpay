package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tariffs struct {
	ID           uint      `gorm:"primaryKey"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string         `gorm:"not null"`
	Description  string         `gorm:"not null"`
	Price        int            `gorm:"not null"`
	Currency     string         `gorm:"not null"`
	ValidityDays int            `gorm:"not null"`
	IsDefault    int            `gorm:"default:0"`
}
