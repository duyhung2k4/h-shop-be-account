package service

import (
	"app/config"
	"app/model"

	"gorm.io/gorm"
)

type loginGoogleService struct {
	db *gorm.DB
}

type LoginGoogleService interface {
	CheckExistUser(userCheck model.User) (bool, error)
}

func (l *loginGoogleService) CheckExistUser(userCheck model.User) (bool, error) {
	var user *model.User

	if err := l.db.Model(&model.User{}).Where("email = ?", userCheck.Email).First(&user).Error; err != nil {
		return false, err
	}

	if user == nil {
		return true, nil
	}

	return false, nil
}

func NewGoginGoogleService() LoginGoogleService {
	return &loginGoogleService{
		db: config.GetDB(),
	}
}
