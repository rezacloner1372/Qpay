package model

import (
	"github.com/jinzhu/gorm"
)

type Roles struct {
	gorm.Model
	Name      string `gorm:"unique_index"`
	GuardName string `gorm:"unique_index"`
}
