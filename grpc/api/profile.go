package api

import (
	"app/config"
	"app/grpc/convert"
	"app/grpc/proto"
	"app/model"
	"context"

	"gorm.io/gorm"
)

type profileGRPC struct {
	db             *gorm.DB
	convertProfile convert.ConvertProfileGRPC
	convertUser    convert.ConvertUser
	convertRole    convert.ConvertRole
	proto.UnsafeProfileServiceServer
}

func (g *profileGRPC) GetProfile(ctx context.Context, pb *proto.GetProfileReq) (*proto.Profile, error) {
	profileID := pb.ProfileID
	var profile *model.Profile

	if err := g.db.
		Model(&model.Profile{}).
		Preload("User").
		Preload("User.Role").
		Where("id = ?", profileID).First(&profile).Error; err != nil {
		return nil, err
	}

	profileGRPC := g.convertProfile.ConvertToGRPC(*profile)
	profileGRPC.User = g.convertUser.ConvertToGRPC(*profile.User)
	profileGRPC.User.Role = g.convertRole.ConvertToGRPC(*profile.User.Role)

	return profileGRPC, nil
}

func NewProfileGRPC() proto.ProfileServiceServer {
	return &profileGRPC{
		db:             config.GetDB(),
		convertProfile: convert.NewConvertProfileGRPC(),
		convertUser:    convert.NewConvertUserGRPC(),
		convertRole:    convert.NewConvertRoleGRPC(),
	}
}
