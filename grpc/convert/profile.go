package convert

import (
	"app/grpc/proto"
	"app/model"

	"google.golang.org/protobuf/types/known/timestamppb"
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
		Birth:     timestamppb.New(profile.Birth),
		Phone:     profile.Phone,
		Email:     profile.Email,
		Picture:   profile.Picture,
		Sub:       profile.Sub,

		CreatedAt: timestamppb.New(profile.CreatedAt),
		UpdatedAt: timestamppb.New(profile.UpdatedAt),
		DeletedAt: timestamppb.New(profile.DeletedAt.Time),
	}
}

func NewConvertProfileGRPC() ConvertProfileGRPC {
	return &convertProfileGRPC{}
}
