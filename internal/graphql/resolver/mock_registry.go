package resolver

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/go-genproto/dictybaseapis/identity"
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"google.golang.org/grpc"
)

func AddAPIEndpoint(key, endpoint string) {
}

// AddAPIClient adds a new entry to the hashmap
func AddAPIConnection(key string, conn *grpc.ClientConn) {
}

// GetAPIClient looks up a client in the hashmap
func GetAPIConnection(key string) (conn *grpc.ClientConn) {
}

func GetUserClient(key string) user.UserServiceClient {
	return user.UserServiceClient
}

func GetRoleClient(key string) user.RoleServiceClient {
	return user.RoleServiceClient
}

func GetPermissionClient(key string) user.PermissionServiceClient {
	return user.PermissionServiceClient
}

func GetStockClient(key string) stock.StockServiceClient {
	return stock.StockServiceClient
}

func GetOrderClient(key string) order.OrderServiceClient {
	return order.OrderServiceClient
}

func GetContentClient(key string) content.ContentServiceClient {
	return content.ContentServiceClient
}

func GetAnnotationClient(key string) annotation.TaggedAnnotationServiceClient {
	return annotation.TaggedAnnotationServiceClient
}

func GetAuthClient(key string) auth.AuthServiceClient {
	return auth.AuthServiceClient
}

func GetIdentityClient(key string) identity.IdentityServiceClient {
	return identity.IdentityServiceClient
}

func GetAPIEndpoint(key string) string {
	return ""
}
