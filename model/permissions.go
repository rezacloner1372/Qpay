package models

import (
	"github.com/jinzhu/gorm"
)

type Permissions struct {
	gorm.Model
	ID        int
	Name      string `gorm:"unique_index"`
	GuardName string `gorm:"unique_index"`
}
