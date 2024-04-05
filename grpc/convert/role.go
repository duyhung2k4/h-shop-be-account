package convert

import (
	"app/grpc/proto"
	"app/model"
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

		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		DeletedAt: user.DeletedAt.Time.Unix(),
	}
}

func NewConvertRoleGRPC() ConvertRole {
	return &convertRole{}
}
