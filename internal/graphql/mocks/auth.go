package mocks

import (
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

// func mockAuth() *auth.Auth {
// 	return &auth.Auth{}
// }

func mockedAuthClient() *clients.AuthServiceClient {
	mockedAuthClient := new(clients.AuthServiceClient)
	return mockedAuthClient
}
