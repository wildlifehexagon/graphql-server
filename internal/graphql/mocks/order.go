package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/mock"
)

var MockOrderAttributes = &order.OrderAttributes{
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
}

var singleMockOrder = &order.Order{
	Data: &order.Order_Data{
		Type:       "order",
		Id:         "999",
		Attributes: MockOrderAttributes,
	},
}

var mockCollection = &order.OrderCollection_Data{
	Type:       "order",
	Id:         "999",
	Attributes: MockOrderAttributes,
}

func mockOrder() *order.Order {
	return singleMockOrder
}

func mockOrderCollection() *order.OrderCollection {
	var orders []*order.OrderCollection_Data
	orders = append(orders, mockCollection, mockCollection, mockCollection)
	return &order.OrderCollection{
		Data: orders,
		Meta: &order.Meta{
			NextCursor: 10000,
			Limit:      10,
			Total:      int64(len(orders)),
		},
	}
}

func mockUpdateOrder() *order.Order {
	return &order.Order{
		Data: &order.Order_Data{
			Type: "order",
			Id:   "999",
			Attributes: &order.OrderAttributes{
				CreatedAt:        ptypes.TimestampNow(),
				UpdatedAt:        ptypes.TimestampNow(),
				Courier:          "FedEx",
				CourierAccount:   "444444",
				Comments:         "Please send ASAP",
				Payment:          "credit",
				PurchaseOrderNum: "987654",
				Status:           1, // growing
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
	).Return(mockOrder(), nil).
		On(
			"ListOrders",
			mock.Anything,
			mock.AnythingOfType("*order.ListParameters"),
		).Return(mockOrderCollection(), nil).
		On(
			"CreateOrder",
			mock.Anything,
			mock.AnythingOfType("*order.NewOrder"),
		).Return(mockOrder(), nil).
		On(
			"UpdateOrder",
			mock.Anything,
			mock.AnythingOfType("*order.OrderUpdate"),
		).Return(mockUpdateOrder(), nil)
	return mockedOrderClient
}
