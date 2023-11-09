package models

import (
	"github.com/jinzhu/gorm"
)

type RolesPermissions struct {
	gorm.Model
	roleId       int
	permissionId int
	GuardName    string `gorm:"unique_index"`
}
