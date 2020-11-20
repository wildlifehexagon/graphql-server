package stock

import (
	"context"
	"testing"

	"github.com/dictyBase/graphql-server/internal/graphql/cache"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/dictyBase/graphql-server/internal/graphql/models"
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

func TestPlasmidGenes(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)
	r := plasmidResolver(mocks.MockedGenModClient())
	rc := r.Registry.GetRedisRepository(cache.RedisKey)
	rc.HSet(cache.GeneHash, "DDB_G0285425", "gpaD")
	g, err := r.Genes(context.Background(), mockPlasmidInput)
	assert.NoError(err, "expect no error from getting associated genes")
	genes := []*models.Gene{}
	genes = append(genes, &models.Gene{
		ID:   "DDB_G0285425",
		Name: "gpaD",
		Goas: nil,
	})
	assert.ElementsMatch(g, genes, "should match associated genes")
}
