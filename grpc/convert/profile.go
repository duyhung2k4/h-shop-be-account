package convert

import (
	"app/grpc/proto"
	"app/model"
)

type convertProfileGRPC struct {
}

type ConvertProfileGRPC interface {
	ConvertToGRPC(model.Profile) *proto.Profile
}

func (c *convertProfileGRPC) ConvertToGRPC(profile model.Profile) *proto.Profile {
	return &proto.Profile{
		ID:        uint64(profile.ID),
		UserID:    uint64(profile.UserID),
		Firstname: profile.Firstname,
		Lastname:  profile.Lastname,
		Name:      profile.Name,
		Address:   profile.Address,
		Gender:    profile.Gender,
		Birth:     profile.Birth.Unix(),
		Phone:     profile.Phone,
		Email:     profile.Email,
		Picture:   profile.Picture,
		Sub:       profile.Sub,

		CreatedAt: profile.CreatedAt.Unix(),
		UpdatedAt: profile.UpdatedAt.Unix(),
		DeletedAt: profile.DeletedAt.Time.Unix(),
	}
}

func NewConvertProfileGRPC() ConvertProfileGRPC {
	return &convertProfileGRPC{}
}
