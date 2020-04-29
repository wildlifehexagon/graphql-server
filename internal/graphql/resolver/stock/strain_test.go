package stock

import (
	"context"
	"testing"

	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/stretchr/testify/assert"
)

func TestSystematicName(t *testing.T) {
	assert := assert.New(t)
	r := &StrainResolver{
		Client:           mocks.MockedStockClient(),
		UserClient:       mocks.MockedUserClient(),
		AnnotationClient: mocks.MockedSysAnnoClient(),
		Registry:         &mocks.MockRegistry{},
		Logger:           mocks.TestLogger(),
	}
	id := "123456"
	sn, err := r.AnnotationClient.GetEntryAnnotation(context.Background(), &annotation.EntryAnnotationRequest{
		Tag:      registry.SysnameTag,
		Ontology: registry.DictyAnnoOntology,
		EntryId:  id,
	})
	assert.NoError(err, "expect no error from getting strain characteristics")
	assert.Exactly(sn.Data.Attributes.Value, mocks.MockSysNameAnno.Data.Attributes.Value, "should match systematic name")
}
