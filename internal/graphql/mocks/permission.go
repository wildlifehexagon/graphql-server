package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

func mockPermission() *user.Permission {
	return &user.Permission{}
}

func mockedPermissionClient() *clients.PermissionServiceClient {
	mockedPermissionClient := new(clients.PermissionServiceClient)
	return mockedPermissionClient
}
