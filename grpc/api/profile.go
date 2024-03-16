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

	return g.convertProfile.ConvertProfile(*profile), nil
}

func NewProfileGRPC() proto.ProfileServiceServer {
	return &profileGRPC{
		db:             config.GetDB(),
		convertProfile: convert.NewConvertProfileGRPC(),
	}
}
