package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID    uint       `json:"userId"`
	Firstname *string    `json:"firstname"`
	Lastname  *string    `json:"lastname"`
	Name      *string    `json:"name"`
	Address   *string    `json:"address"`
	Gender    *string    `json:"gender"`
	Birth     *time.Time `json:"birth"`
	Phone     *string    `json:"phone"`
	Email     *string    `json:"email" gorm:"unique"`
	Picture   *string    `json:"picture"`
	Sub       *string    `json:"sub" gorm:"unique"`

	User *User `json:"user" gorm:"foreignKey:UserID"`
}
