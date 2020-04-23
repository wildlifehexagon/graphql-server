package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/golang/protobuf/ptypes"
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

func mockedOrderClient() *clients.OrderServiceClient {
	mockedOrderClient := new(clients.OrderServiceClient)
	mockedOrderClient.On(
		"GetOrder",
		mock.Anything,
		mock.AnythingOfType("*order.OrderId"),
	).Return(mockOrder(), nil)
	return mockedOrderClient
}
