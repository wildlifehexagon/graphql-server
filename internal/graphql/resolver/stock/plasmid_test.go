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

func plasmidResolver(annoClient *clients.TaggedAnnotationServiceClient) *PlasmidResolver {
	return &PlasmidResolver{
		Client:           mocks.MockedStockClient(),
		UserClient:       mocks.MockedUserClient(),
		AnnotationClient: annoClient,
		Registry:         &mocks.MockRegistry{},
		Logger:           mocks.TestLogger(),
	}
}

var mockPlasmidInput = &models.Plasmid{
	Data: &stock.Plasmid_Data{
		Type:       "plasmid",
		Id:         "DBP0000120",
		Attributes: mocks.MockPlasmidAttributes,
	},
}

func TestPlasmidID(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedAnnotationClient())
	p, err := r.ID(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting plasmid id")
	assert.Exactly(p, mockPlasmidInput.Data.Id, "should match id")
}

func TestPlasmidSummary(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedAnnotationClient())
	p, err := r.Summary(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting summary")
	assert.Exactly(p, &mockPlasmidInput.Data.Attributes.Summary, "should match summary")
}

func TestPlasmidEditableSummary(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedAnnotationClient())
	p, err := r.EditableSummary(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting editable summary")
	assert.Exactly(p, &mockPlasmidInput.Data.Attributes.EditableSummary, "should match editable summary")
}

func TestPlasmidDepositor(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedAnnotationClient())
	p, err := r.Depositor(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting depositor")
	assert.Exactly(p, mockPlasmidInput.Data.Attributes.Depositor, "should match depositor")
}

func TestSequence(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedAnnotationClient())
	p, err := r.Sequence(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting plasmid sequence")
	assert.Exactly(p, &mockPlasmidInput.Data.Attributes.Sequence, "should match sequence")
}

func TestName(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedAnnotationClient())
	p, err := r.Name(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting plasmid name")
	assert.Exactly(p, mockPlasmidInput.Data.Attributes.Name, "should match name")
}

func TestPlasmidInStock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedInStockClient())
	p, err := r.InStock(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting plasmid inventory")
	assert.True(p, "should return true after finding inventory")
}
