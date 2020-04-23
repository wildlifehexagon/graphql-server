package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

func mockRole() *user.Role {
	return &user.Role{}
}

func mockedRoleClient() *clients.RoleServiceClient {
	mockedRoleClient := new(clients.RoleServiceClient)
	return mockedRoleClient
}
