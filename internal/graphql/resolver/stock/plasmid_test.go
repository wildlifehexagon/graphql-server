package stock

import (
	"context"
	"testing"

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

var mockPlasmidInput = ConvertToPlasmidModel("DBP0000120", mocks.MockPlasmidAttributes)

func TestPlasmidInStock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedInStockClient())
	p, err := r.InStock(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting plasmid inventory")
	assert.True(p, "should return true after finding inventory")
}
