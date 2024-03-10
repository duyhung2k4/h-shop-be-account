package service

import (
	"app/config"
	"app/dto/request"
	"app/model"
	"errors"

	"gorm.io/gorm"
)

type loginGoogleService struct {
	db *gorm.DB
}

type LoginGoogleService interface {
	CheckExistUser(userCheck request.LoginGoogleRequest) (bool, *model.User, error)
	CreateProfile(userRequest request.LoginGoogleRequest, role model.ROLE) (*model.Profile, error)
	GetProfile(profileId uint) (*model.Profile, error)
}

func (l *loginGoogleService) CheckExistUser(userCheck request.LoginGoogleRequest) (bool, *model.User, error) {
	var user *model.User

	if err := l.db.
		Model(&model.User{}).
		Preload("Profile").
		Preload("Profile.User").
		Preload("Profile.User.Role").
		Where("email = ?", userCheck.Email).
		First(&user).Error; err != nil && err.Error() != "record not found" {
		return false, nil, err
	}

	if user.ID != 0 {
		return true, user, nil
	}

	return false, nil, nil
}

func (l *loginGoogleService) CreateProfile(userRequest request.LoginGoogleRequest, role model.ROLE) (*model.Profile, error) {
	var newProfile model.Profile

	err := l.db.Transaction(func(tx *gorm.DB) error {
		var getRole *model.Role
		if err := l.db.Model(&model.Role{}).Where("Code = ?", role).First(&getRole).Error; err != nil && err.Error() != "record not found" {
			return err
		}
		if getRole == nil {
			return errors.New("role not exist")
		}

		var newUser = model.User{
			RoleID: getRole.ID,
			Email:  userRequest.Email,
		}
		if err := l.db.Model(&model.User{}).Create(&newUser).Error; err != nil {
			tx.Rollback()
			return err
		}

		newProfile = model.Profile{
			UserID:    newUser.ID,
			Firstname: &userRequest.GivenName,
			Lastname:  &userRequest.FamilyName,
			Name:      &userRequest.Name,
			Email:     &userRequest.Email,
			Picture:   &userRequest.Picture,
			Sub:       &userRequest.Sub,
		}
		if err := l.db.Model(&model.Profile{}).Create(&newProfile).Error; err != nil {
			tx.Rollback()
			return err
		}

		newUser.Role = getRole
		newProfile.User = &newUser

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &newProfile, nil
}

func (l *loginGoogleService) GetProfile(profileId uint) (*model.Profile, error) {
	var profile *model.Profile

	if err := l.db.
		Model(&model.Profile{}).
		Preload("User").
		Preload("User.Role").
		Where("id = ?", profileId).
		First(&profile).Error; err != nil {
		return nil, err
	}

	return profile, nil
}

func NewGoginGoogleService() LoginGoogleService {
	return &loginGoogleService{
		db: config.GetDB(),
	}
}
