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

var MockGenModAnno = &annotation.TaggedAnnotation{
	Data: &annotation.TaggedAnnotation_Data{
		Type: "annotation",
		Id:   "123456",
		Attributes: &annotation.TaggedAnnotationAttributes{
			Value:     "exogenous mutation",
			EntryId:   "DBS0236922",
			CreatedBy: "dsc@dictycr.org",
			CreatedAt: ptypes.TimestampNow(),
			Tag:       registry.MuttypeTag,
			Ontology:  registry.DictyAnnoOntology,
			Version:   1,
		},
	},
}

var MockMutMethodAnno = &annotation.TaggedAnnotation{
	Data: &annotation.TaggedAnnotation_Data{
		Type: "annotation",
		Id:   "123456",
		Attributes: &annotation.TaggedAnnotationAttributes{
			Value:     "Random Insertion",
			EntryId:   "DBS0236922",
			CreatedBy: "dsc@dictycr.org",
			CreatedAt: ptypes.TimestampNow(),
			Tag:       registry.MutmethodTag,
			Ontology:  registry.DictyAnnoOntology,
			Version:   1,
		},
	},
}

var MockGenotypeAnno = &annotation.TaggedAnnotation{
	Data: &annotation.TaggedAnnotation_Data{
		Type: "annotation",
		Id:   "123456",
		Attributes: &annotation.TaggedAnnotationAttributes{
			Value:     "axeA1,axeB1,axeC1,sadA-[sadA-KO],[pSadA-GFP],bsR,neoR",
			EntryId:   "DBS0236922",
			CreatedBy: "dsc@dictycr.org",
			CreatedAt: ptypes.TimestampNow(),
			Tag:       registry.GenoTag,
			Ontology:  registry.DictyAnnoOntology,
			Version:   1,
		},
	},
}

func MockedAnnotationClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	return mockedAnnoClient
}

func MockedSysNameAnnoClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	mockedAnnoClient.On(
		"GetEntryAnnotation",
		mock.Anything,
		mock.AnythingOfType("*annotation.EntryAnnotationRequest"),
	).Return(MockSysNameAnno, nil)
	return mockedAnnoClient
}

func MockedGenModClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	mockedAnnoClient.On(
		"GetEntryAnnotation",
		mock.Anything,
		mock.AnythingOfType("*annotation.EntryAnnotationRequest"),
	).Return(MockGenModAnno, nil)
	return mockedAnnoClient
}

func MockedMutMethodClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	mockedAnnoClient.On(
		"GetEntryAnnotation",
		mock.Anything,
		mock.AnythingOfType("*annotation.EntryAnnotationRequest"),
	).Return(MockMutMethodAnno, nil)
	return mockedAnnoClient
}

func MockedGenotypeClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	mockedAnnoClient.On(
		"GetEntryAnnotation",
		mock.Anything,
		mock.AnythingOfType("*annotation.EntryAnnotationRequest"),
	).Return(MockGenotypeAnno, nil)
	return mockedAnnoClient
}
