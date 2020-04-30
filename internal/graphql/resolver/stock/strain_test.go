package stock

import (
	"context"
	"testing"

	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
	"github.com/golang/protobuf/ptypes"

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

func TestID(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	sn, err := r.ID(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting strain id")
	assert.Exactly(sn, mockStrainInput.Data.Id, "should match id")
}

func TestCreatedAt(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	ca, err := r.CreatedAt(context.Background(), mockStrainInput)
	timestamp, _ := ptypes.Timestamp(mockStrainInput.Data.Attributes.CreatedAt)
	assert.NoError(err, "expect no error from getting created_at")
	assert.Exactly(ca, &timestamp, "should match created_at")
}

func TestUpdatedAt(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	ua, err := r.UpdatedAt(context.Background(), mockStrainInput)
	timestamp, _ := ptypes.Timestamp(mockStrainInput.Data.Attributes.UpdatedAt)
	assert.NoError(err, "expect no error from getting updated_at")
	assert.Exactly(ua, &timestamp, "should match updated_at")
}

func TestSummary(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	sn, err := r.Summary(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting summary")
	assert.Exactly(sn, &mockStrainInput.Data.Attributes.Summary, "should match summary")
}

func TestEditableSummary(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	sn, err := r.EditableSummary(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting editable summary")
	assert.Exactly(sn, &mockStrainInput.Data.Attributes.EditableSummary, "should match editable summary")
}

func TestDepositor(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	sn, err := r.Depositor(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting depositor")
	assert.Exactly(sn, mockStrainInput.Data.Attributes.Depositor, "should match depositor")
}

func TestLabel(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	sn, err := r.Label(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting label")
	assert.Exactly(sn, mockStrainInput.Data.Attributes.Label, "should match label")
}

func TestSpecies(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	sn, err := r.Species(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting species")
	assert.Exactly(sn, mockStrainInput.Data.Attributes.Species, "should match species")
}

func TestPlasmid(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := strainResolver(mocks.MockedAnnotationClient())
	sn, err := r.Plasmid(context.Background(), mockStrainInput)
	assert.NoError(err, "expect no error from getting plasmid")
	assert.Exactly(sn, &mockStrainInput.Data.Attributes.Plasmid, "should match plasmid")
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
	assert.ElementsMatch(g, gl, "should match genotypes")
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
