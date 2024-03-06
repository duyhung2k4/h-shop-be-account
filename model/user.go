package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID   uint   `json:"roleId"`
	Username string `json:"username"`
	Email    string `json:"email"`

	Role    *Role    `json:"role" gorm:"foreignKey:RoleID"`
	Profile *Profile `json:"profile" gorm:"foreignKey:UserID"`
}
