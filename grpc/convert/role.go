package convert

import (
	"app/grpc/proto"
	"app/model"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type convertRole struct{}

type ConvertRole interface {
	ConvertToGRPC(model.Role) *proto.Role
}

func (c *convertRole) ConvertToGRPC(user model.Role) *proto.Role {
	return &proto.Role{
		ID:   uint64(user.ID),
		Code: string(user.Code),
		Name: user.Name,

		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		DeletedAt: timestamppb.New(user.DeletedAt.Time),
	}
}

func NewConvertRoleGRPC() ConvertRole {
	return &convertRole{}
}
