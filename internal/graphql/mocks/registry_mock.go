package mocks

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/go-genproto/dictybaseapis/identity"
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/mocks/clients"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
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

type MockRegistry struct{}

func (mr *MockRegistry) AddAPIEndpoint(key, endpoint string) {
}

// AddAPIClient adds a new entry to the hashmap
func (mr *MockRegistry) AddAPIConnection(key string, conn *grpc.ClientConn) {
}

// GetAPIClient looks up a client in the hashmap
func (mr *MockRegistry) GetAPIConnection(key string) (conn *grpc.ClientConn) {
	return &grpc.ClientConn{}
}

func (mr *MockRegistry) GetUserClient(key string) user.UserServiceClient {
	return user.NewUserServiceClient(mr.GetAPIConnection(key))
}

func (mr *MockRegistry) GetRoleClient(key string) user.RoleServiceClient {
	return user.NewRoleServiceClient(mr.GetAPIConnection(key))
}

func (mr *MockRegistry) GetPermissionClient(key string) user.PermissionServiceClient {
	return user.NewPermissionServiceClient(mr.GetAPIConnection(key))
}

func (mr *MockRegistry) GetStockClient(key string) stock.StockServiceClient {
	return stock.NewStockServiceClient(mr.GetAPIConnection(key))
}

func (mr *MockRegistry) GetOrderClient(key string) order.OrderServiceClient {
	return mockedOrderClient()
}

func (mr *MockRegistry) GetContentClient(key string) content.ContentServiceClient {
	return content.NewContentServiceClient(mr.GetAPIConnection(key))
}

func (mr *MockRegistry) GetAnnotationClient(key string) annotation.TaggedAnnotationServiceClient {
	return annotation.NewTaggedAnnotationServiceClient(mr.GetAPIConnection(key))
}

func (mr *MockRegistry) GetAuthClient(key string) auth.AuthServiceClient {
	return auth.NewAuthServiceClient(mr.GetAPIConnection(key))
}

func (mr *MockRegistry) GetIdentityClient(key string) identity.IdentityServiceClient {
	return identity.NewIdentityServiceClient(mr.GetAPIConnection(key))
}

func (mr MockRegistry) GetAPIEndpoint(key string) string {
	return key
}
