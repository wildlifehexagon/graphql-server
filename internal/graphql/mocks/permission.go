package mocks

import (
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

// func mockPermission() *user.Permission {
// 	return &user.Permission{}
// }

func MockedPermissionClient() *clients.PermissionServiceClient {
	mockedPermissionClient := new(clients.PermissionServiceClient)
	return mockedPermissionClient
}
