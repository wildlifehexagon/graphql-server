package mocks

import (
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

// func mockIdentity() *identity.Identity {
// 	return &identity.Identity{}
// }

func MockedIdentityClient() *clients.IdentityServiceClient {
	mockedIdentityClient := new(clients.IdentityServiceClient)
	return mockedIdentityClient
}
