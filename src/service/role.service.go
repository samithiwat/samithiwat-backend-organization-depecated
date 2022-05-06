package service

import "github.com/samithiwat/samithiwat-backend-organization/src/proto"

type RoleClient interface {
	FindOne(*proto.Role) []error
	Create(*proto.Role) []error
	Update(*proto.Role) []error
	Delete(*proto.Role) []error
}

type RoleService struct {
	client *RoleClient
}

func NewRoleService(client *RoleClient) *RoleService {
	return &RoleService{
		client: client,
	}
}
