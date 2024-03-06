package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	RoleID   uint   `json:"roleId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`

	Role    *Role    `json:"role" gorm:"foreignKey:RoleID"`
	Profile *Profile `json:"profile" gorm:"foreignKey:UserID"`
}
