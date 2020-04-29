package stock

import (
	"context"
	"testing"

	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/graphql-server/internal/graphql/models"

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/stretchr/testify/assert"
)

var mockInput = &models.Strain{
	Data: &stock.Strain_Data{
		Type:       "strain",
		Id:         "DBS0236922",
		Attributes: mocks.MockStrainAttributes,
	},
}

func TestSystematicName(t *testing.T) {
	assert := assert.New(t)
	r := &StrainResolver{
		Client:           mocks.MockedStockClient(),
		UserClient:       mocks.MockedUserClient(),
		AnnotationClient: mocks.MockedSysAnnoClient(),
		Registry:         &mocks.MockRegistry{},
		Logger:           mocks.TestLogger(),
	}
	sn, err := r.SystematicName(context.Background(), mockInput)
	assert.NoError(err, "expect no error from getting systematic name")
	assert.Exactly(sn, mocks.MockSysNameAnno.Data.Attributes.Value, "should match systematic name")
}
