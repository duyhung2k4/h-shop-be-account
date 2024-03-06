package model

import (
	"gorm.io/gorm"
)

type ROLE string

const (
	USER    ROLE = "user"
	ADMIN   ROLE = "admin"
	SHIPPER ROLE = "shipper"
)

type Role struct {
	gorm.Model
	Code ROLE   `json:"code" gorm:"unique"`
	Name string `json:"name"`

	Users []User `json:"users" gorm:"foreignKey:RoleID"`
}
