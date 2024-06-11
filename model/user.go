package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID   uint   `json:"roleId" gorm:"uniqueIndex:idx_role_email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"uniqueIndex:idx_role_email"`

	Role    *Role    `json:"role" gorm:"foreignKey:RoleID"`
	Profile *Profile `json:"profile" gorm:"foreignKey:UserID"`
}
