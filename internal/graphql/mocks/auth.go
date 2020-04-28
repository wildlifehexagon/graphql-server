package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/go-genproto/dictybaseapis/identity"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/mock"
)

func mockAuth() *auth.Auth {
	return &auth.Auth{
		Token:        "test token",
		RefreshToken: "test refresh token",
		Identity: &identity.Identity{
			Data: &identity.IdentityData{
				Type: "identity",
				Id:   123,
				Attributes: &identity.IdentityAttributes{
					Identifier: "art@vandelay.com",
					Provider:   "google",
					UserId:     999,
					CreatedAt:  ptypes.TimestampNow(),
					UpdatedAt:  ptypes.TimestampNow(),
				},
			},
		},
		User: &user.User{
			Data: &user.UserData{
				Type: "user",
				Id:   999,
				Attributes: &user.UserAttributes{
					FirstName: "Art",
					LastName:  "Vandelay",
					Email:     "art@vandelay.com",
					IsActive:  true,
					CreatedAt: ptypes.TimestampNow(),
					UpdatedAt: ptypes.TimestampNow(),
				},
			},
		},
	}
}

func MockedAuthClient() *clients.AuthServiceClient {
	mockedAuthClient := new(clients.AuthServiceClient)
	mockedAuthClient.On(
		"Login",
		mock.AnythingOfType("*context.valueCtx"),
		mock.AnythingOfType("*auth.NewLogin"),
	).Return(mockAuth(), nil)
	return mockedAuthClient
}
