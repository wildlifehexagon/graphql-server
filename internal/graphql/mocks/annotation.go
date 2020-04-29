package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/mock"
)

var MockSysNameAnno = &annotation.TaggedAnnotation{
	Data: &annotation.TaggedAnnotation_Data{
		Type: "annotation",
		Id:   "123456",
		Attributes: &annotation.TaggedAnnotationAttributes{
			Value:     "DBS0236922",
			EntryId:   "DBS0236922",
			CreatedBy: "dsc@dictycr.org",
			CreatedAt: ptypes.TimestampNow(),
			Tag:       registry.SysnameTag,
			Ontology:  registry.DictyAnnoOntology,
			Version:   1,
		},
	},
}

func MockedAnnotationClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	return mockedAnnoClient
}

func MockedSysAnnoClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	mockedAnnoClient.On(
		"GetEntryAnnotation",
		mock.Anything,
		mock.AnythingOfType("*annotation.EntryAnnotationRequest"),
	).Return(MockSysNameAnno, nil)
	return mockedAnnoClient
}
