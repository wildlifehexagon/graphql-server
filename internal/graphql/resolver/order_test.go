package resolver

import (
	"context"
	"testing"

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	assert := assert.New(t)
	ord := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	_, err := ord.Order(context.Background(), "999")
	assert.NoError(err, "expect no error from getting order information")
	// assert.Exactly(name, "DBS0236922", "should match systematic name")
}
