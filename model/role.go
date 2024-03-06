package model

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Code string `json:"code" gorm:"unique"`
	Name string `json:"name"`

	Users []User `json:"users" gorm:"foreignKey:RoleID"`
}
