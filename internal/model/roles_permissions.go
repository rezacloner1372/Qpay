package model

import (
	"github.com/jinzhu/gorm"
)

type RolesPermissions struct {
	gorm.Model
	roleId       int
	permissionId int
	GuardName    string      `gorm:"unique_index"`
	Roles        Roles       `gorm:"references:role_id"`
	Permissions  Permissions `gorm:"references:permission_id"`
}
