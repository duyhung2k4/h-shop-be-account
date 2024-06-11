package service

import (
	"app/config"
	"app/dto/request"
	"app/grpc/proto"
	"app/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type loginGoogleService struct {
	db             *gorm.DB
	clientShopGRPC proto.ShopServiceClient
}

type LoginGoogleService interface {
	CheckExistUser(userCheck request.LoginGoogleRequest) (bool, *model.User, error)
	CreateProfile(userRequest request.LoginGoogleRequest) (*model.Profile, error)
	GetProfile(profileId uint) (*model.Profile, error)
}

func (s *loginGoogleService) CheckExistUser(userCheck request.LoginGoogleRequest) (bool, *model.User, error) {
	var user *model.User
	var role *model.Role

	s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Role{}).Where("code = ?", userCheck.Role).Find(&role).Error; err != nil {
			return err
		}

		if err := tx.
			Model(&model.User{}).
			Preload("Profile").
			Preload("Profile.User").
			Preload("Profile.User.Role").
			Where("email = ? AND role_id = ?", userCheck.Email, role.ID).
			First(&user).Error; err != nil && err.Error() != "record not found" {
			return nil
		}
		return nil
	})

	if user.ID != 0 {
		return true, user, nil
	}

	return false, nil, nil
}

func (s *loginGoogleService) CreateProfile(userRequest request.LoginGoogleRequest) (*model.Profile, error) {
	var newProfile model.Profile

	err := s.db.Transaction(func(tx *gorm.DB) error {
		var getRole *model.Role
		if err := s.db.Model(&model.Role{}).Where("Code = ?", userRequest.Role).First(&getRole).Error; err != nil && err.Error() != "record not found" {
			return err
		}
		if getRole == nil {
			return errors.New("role not exist")
		}

		var newUser = model.User{
			RoleID: getRole.ID,
			Email:  userRequest.Email,
		}
		if err := s.db.Model(&model.User{}).Create(&newUser).Error; err != nil {
			tx.Rollback()
			return err
		}

		newProfile = model.Profile{
			UserID:    newUser.ID,
			Firstname: userRequest.GivenName,
			Lastname:  userRequest.FamilyName,
			Name:      userRequest.Name,
			Email:     userRequest.Email,
			Picture:   userRequest.Picture,
			Sub:       userRequest.Sub,
		}
		if err := s.db.Model(&model.Profile{}).Create(&newProfile).Error; err != nil {
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

	var errHandlerGRPC error
	switch role := userRequest.Role; role {
	case model.OWNER_SHOP:
		_, err := s.clientShopGRPC.CreateShop(context.Background(), &proto.CreateShopReq{ProfileId: uint64(newProfile.ID)})
		errHandlerGRPC = err
	}

	if errHandlerGRPC != nil {
		return nil, err
	}

	return &newProfile, nil
}

func (s *loginGoogleService) GetProfile(profileId uint) (*model.Profile, error) {
	var profile *model.Profile

	if err := s.db.
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
		db:             config.GetDB(),
		clientShopGRPC: proto.NewShopServiceClient(config.GetClientGRPCShop()),
	}
}
