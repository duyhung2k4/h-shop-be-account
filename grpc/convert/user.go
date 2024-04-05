package convert

import (
	"app/grpc/proto"
	"app/model"
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

		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		DeletedAt: user.DeletedAt.Time.Unix(),
	}
}

func NewConvertUserGRPC() ConvertUser {
	return &convertUser{}
}
