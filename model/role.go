package model

import (
	"gorm.io/gorm"
)

type ROLE string

const (
	USER            ROLE = "user"
	SHIPPER         ROLE = "shipper"
	OWNER_SHOP      ROLE = "owner_shop"
	OWNER_WAREHOUSE ROLE = "owner_warehouse"
)

type Role struct {
	gorm.Model
	Code ROLE   `json:"code" gorm:"unique"`
	Name string `json:"name"`

	Users []User `json:"users" gorm:"foreignKey:RoleID"`
}
