package stock

import (
	"context"
	"testing"

	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/models"

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/stretchr/testify/assert"
)

func strainResolver(annoClient *clients.TaggedAnnotationServiceClient) *StrainResolver {
	return &StrainResolver{
		Client:           mocks.MockedStockClient(),
		UserClient:       mocks.MockedUserClient(),
		AnnotationClient: annoClient,
		Registry:         &mocks.MockRegistry{},
		Logger:           mocks.TestLogger(),
	}
}

var mockStrainInput = &models.Strain{
	Data: &stock.Strain_Data{
		Type:       "strain",
		Id:         "DBS0236922",
		Attributes: mocks.MockStrainAttributes,
	},
}

func TestSystematicName(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedSysNameAnnoClient())
	sn, err := r.SystematicName(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting systematic name")
	assert.Exactly(sn, mocks.MockSysNameAnno.Data.Attributes.Value, "should match systematic name")
}

func TestGeneticModification(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedGenModClient())
	g, err := r.GeneticModification(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting genetic modification")
	assert.Exactly(g, &mocks.MockGenModAnno.Data.Attributes.Value, "should match genetic modification")
}

func TestMutagenesisMethod(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedMutMethodClient())
	m, err := r.MutagenesisMethod(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting mutagenesis method")
	assert.Exactly(m, &mocks.MockMutMethodAnno.Data.Attributes.Value, "should match mutagenesis method")
}

func TestGenotypes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedMutMethodClient())
	g, err := r.Genotypes(context.Background(), mockStrainInput)
	gl := []*string{}
	gl = append(gl, &mocks.MockMutMethodAnno.Data.Attributes.Value)
	assert.NoError(err, "expect no error from getting genotypes")
	assert.Exactly(g, gl, "should match genotypes")
}

func TestStrainInStock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedInStockClient())
	g, err := r.InStock(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting strain inventory")
	assert.True(g, "should return true after finding inventory")
}
