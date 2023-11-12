package model

import (
	"github.com/jinzhu/gorm"
)

type Permissions struct {
	gorm.Model
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"unique_index"`
	GuardName string `gorm:"unique_index"`
}
