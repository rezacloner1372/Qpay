package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Tariffs struct {
	gorm.Model
	Name         string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	Price        string    `gorm:"not null"`
	Currency     string    `gorm:"not null"`
	ValidityDays int       `gorm:"not null"`
	isDefault    int       `gorm:"default:0"`
	CreationDate time.Time `gorm:"default:CURRENT_TIMESTAMP()"`
}
