package resolver

import (
	"context"
	"testing"

	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func mockOrder() *order.Order {
	return &order.Order{
		Data: &order.Order_Data{
			Type: "order",
			Id:   "999",
			Attributes: &order.OrderAttributes{
				CreatedAt:        ptypes.TimestampNow(),
				UpdatedAt:        ptypes.TimestampNow(),
				Courier:          "USPS",
				CourierAccount:   "123456",
				Comments:         "first order",
				Payment:          "credit",
				PurchaseOrderNum: "987654",
				Status:           0, // in preparation
				Consumer:         "art@vandelayindustries.com",
				Payer:            "george@costanza.com",
				Purchaser:        "thatsgold@jerry.org",
				Items:            []string{"DBS123456"},
			},
		},
	}
}

func mockedOrderClient() *mocks.OrderServiceClient {
	mockedOrderClient := new(mocks.OrderServiceClient)
	mockedOrderClient.On(
		"GetOrder",
		mock.Anything,
		mock.AnythingOfType("*order.OrderId"),
	).Return(mockOrder(), nil)
	return mockedOrderClient
}

func TestOrder(t *testing.T) {
	assert := assert.New(t)
	ord := &QueryResolver{}
	_, err := ord.Order(context.Background(), "999")
	assert.NoError(err, "expect no error from getting order information")
	// assert.Exactly(name, "DBS0236922", "should match systematic name")
}
