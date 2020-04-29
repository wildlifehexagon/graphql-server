package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/mock"
)

func MockTagAnno(value, tag string) *annotation.TaggedAnnotation {
	return &annotation.TaggedAnnotation{
		Data: &annotation.TaggedAnnotation_Data{
			Type: "annotation",
			Id:   "123456",
			Attributes: &annotation.TaggedAnnotationAttributes{
				Value:     value,
				EntryId:   "DBS0236922",
				CreatedBy: "dsc@dictycr.org",
				CreatedAt: ptypes.TimestampNow(),
				Tag:       tag,
				Ontology:  registry.DictyAnnoOntology,
				Version:   1,
			},
		},
	}
}

func MockInStockAnno() *annotation.TaggedAnnotationGroupCollection {
	gcdata := []*annotation.TaggedAnnotationGroupCollection_Data{}
	gdata := []*annotation.TaggedAnnotationGroup_Data{}
	gdata = append(gdata, &annotation.TaggedAnnotationGroup_Data{
		Type: "annotation",
		Id:   "489483843",
		Attributes: &annotation.TaggedAnnotationAttributes{
			Version:   1,
			EntryId:   "DBS0235559",
			CreatedBy: "art@vandelay.org",
			CreatedAt: ptypes.TimestampNow(),
			Ontology:  registry.DictyAnnoOntology,
			Tag:       registry.InvLocationTag,
			Value:     "2-9(55-57)",
		},
	})
	gcdata = append(gcdata, &annotation.TaggedAnnotationGroupCollection_Data{
		Type: "annotation_group",
		Group: &annotation.TaggedAnnotationGroup{
			Data:      gdata,
			GroupId:   "4924132",
			CreatedAt: ptypes.TimestampNow(),
			UpdatedAt: ptypes.TimestampNow(),
		},
	})
	return &annotation.TaggedAnnotationGroupCollection{
		Data: gcdata,
	}
}

var MockSysNameAnno = MockTagAnno("DBS0236922", registry.SysnameTag)
var MockGenModAnno = MockTagAnno("exogenous mutation", registry.MuttypeTag)
var MockMutMethodAnno = MockTagAnno("Random Insertion", registry.MutmethodTag)
var MockGenotypeAnno = MockTagAnno("axeA1,axeB1,axeC1,sadA-[sadA-KO],[pSadA-GFP],bsR,neoR", registry.GenoTag)

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

func MockedInStockClient() *clients.TaggedAnnotationServiceClient {
	mockedAnnoClient := new(clients.TaggedAnnotationServiceClient)
	mockedAnnoClient.On(
		"ListAnnotationGroups",
		mock.Anything,
		mock.AnythingOfType("*annotation.ListGroupParameters"),
	).Return(MockInStockAnno(), nil)
	return mockedAnnoClient
}
