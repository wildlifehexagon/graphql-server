package mocks

import (
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

// func mockUser() *user.User {
// 	return &user.User{}
// }

func mockedUserClient() *clients.UserServiceClient {
	mockedUserClient := new(clients.UserServiceClient)
	return mockedUserClient
}
