package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tariffs struct {
	gorm.Model
	ID           uint      `gorm:"primary_key"`
	Name         string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	Price        string    `gorm:"not null"`
	Currency     string    `gorm:"not null"`
	ValidityDays int       `gorm:"not null"`
	isDefault    int       `gorm:"not null"`
	CreationDate time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
}
