package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

func mockContent() *content.Content {
	return &content.Content{}
}

func mockedContentClient() *clients.ContentServiceClient {
	mockedContentClient := new(clients.ContentServiceClient)
	return mockedContentClient
}
