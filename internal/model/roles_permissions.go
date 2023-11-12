package model

import (
	"github.com/jinzhu/gorm"
)

type RolesPermissions struct {
	gorm.Model
	roleId       uint   `gorm:"column:role_id"`
	permissionId uint   `gorm:"column:permission_id"`
	GuardName    string `gorm:"unique_index"`
}
