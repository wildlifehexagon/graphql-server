package mocks

import (
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
)

// func mockAnnotation() *annotation.TaggedAnnotation {
// 	return &annotation.TaggedAnnotation{}
// }

func mockedAnnotationClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnotationClient := new(clients.TaggedAnnotationServiceClient)
	return mockedAnnotationClient
}
