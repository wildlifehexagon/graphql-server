package mocks

import (
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

// func mockRole() *user.Role {
// 	return &user.Role{}
// }

func MockedRoleClient() *clients.RoleServiceClient {
	mockedRoleClient := new(clients.RoleServiceClient)
	return mockedRoleClient
}
