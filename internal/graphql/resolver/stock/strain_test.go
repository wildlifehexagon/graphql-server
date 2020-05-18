package stock

import (
	"context"
	"testing"

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

var mockStrainInput = ConvertToStrainModel("DBS0236922", mocks.MockStrainAttributes)

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
	assert.ElementsMatch(g, gl, "should match genotypes")
}

func TestNames(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedNamesClient())
	n, err := r.Names(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting names")
	assert.Exactly(n[0], &mocks.MockNamesAnno().Data[0].Attributes.Value, "should match first names value")
	assert.Exactly(n[1], &mocks.MockNamesAnno().Data[1].Attributes.Value, "should match second names value")
}

func TestCharacteristics(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedCharacteristicsClient())
	c, err := r.Characteristics(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting characteristics")
	assert.Exactly(c[0], &mocks.MockCharacteristicsAnno().Data[0].Attributes.Tag, "should match first characteristics value")
	assert.Exactly(c[1], &mocks.MockCharacteristicsAnno().Data[1].Attributes.Tag, "should match second characteristics value")
}

func TestPhenotypes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedPhenotypeClient())
	p, err := r.Phenotypes(context.Background(), mockStrainInput)
	pd := mocks.MockPhenotypeAnno().Data[0]
	assert.NoError(err, "expect no error from getting phenotypes")
	for _, n := range p {
		assert.Exactly(n.Phenotype, pd.Group.Data[0].Attributes.Tag, "should match phenotype")
		assert.Exactly(n.Assay, &pd.Group.Data[1].Attributes.Tag, "should match assay")
		assert.Exactly(n.Environment, &pd.Group.Data[2].Attributes.Tag, "should match environment")
		assert.Exactly(n.Note, &pd.Group.Data[3].Attributes.Value, "should match note")
	}
}

func TestStrainInStock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedInStockClient())
	g, err := r.InStock(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting strain inventory")
	assert.True(g, "should return true after finding inventory")
}
