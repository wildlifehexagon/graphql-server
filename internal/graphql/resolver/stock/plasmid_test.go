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

func TestPlasmidInStock(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedInStockClient())
	g, err := r.InStock(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting plasmid inventory")
	assert.True(g, "should return true after finding inventory")
}
