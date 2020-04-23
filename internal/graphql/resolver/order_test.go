package resolver

import (
	"context"
	"testing"

	"github.com/dictyBase/go-genproto/dictybaseapis/order"

	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/stretchr/testify/assert"
)

func TestOrder(t *testing.T) {
	assert := assert.New(t)
	ord := &QueryResolver{
		Registry: &mocks.MockRegistry{},
		Logger:   mocks.TestLogger(),
	}
	o, err := ord.Order(context.Background(), "999")
	assert.NoError(err, "expect no error from getting order information")
	assert.Exactly(o.Data.Id, "999", "should match id")
	assert.Exactly(o.Data.Attributes.Courier, "USPS", "should match courier")
	assert.Exactly(o.Data.Attributes.CourierAccount, "123456", "should match courier account")
	assert.Exactly(o.Data.Attributes.Comments, "first order", "should match comments")
	assert.Exactly(o.Data.Attributes.Payment, "credit", "should match payment")
	assert.Exactly(o.Data.Attributes.PurchaseOrderNum, "987654", "should match purchase order number")
	assert.Exactly(o.Data.Attributes.Status, order.OrderStatus_In_preparation, "should match status")
	assert.Exactly(o.Data.Attributes.Consumer, "art@vandelayindustries.com", "should match consumer")
	assert.Exactly(o.Data.Attributes.Payer, "george@costanza.com", "should match payer")
	assert.Exactly(o.Data.Attributes.Purchaser, "thatsgold@jerry.org", "should match purchaser")
	assert.ElementsMatch(o.Data.Attributes.Items, []string{"DBS123456"}, "should match items")
}
