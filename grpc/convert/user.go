package convert

import (
	"app/grpc/proto"
	"app/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type convertUser struct{}

type ConvertUser interface {
	ConvertToGRPC(model.User) *proto.User
}

func (c *convertUser) ConvertToGRPC(user model.User) *proto.User {
	return &proto.User{
		ID:       uint64(user.ID),
		RoleID:   uint64(user.RoleID),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,

		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		DeletedAt: timestamppb.New(user.DeletedAt.Time),
	}
}

func NewConvertUserGRPC() ConvertUser {
	return &convertUser{}
}
