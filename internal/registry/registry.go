package registry

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/emirpasic/gods/maps/hashmap"
	"google.golang.org/grpc"
)

const (
	USER        = "user"
	ROLE        = "role"
	PERMISSION  = "permission"
	PUBLICATION = "publication"
	STOCK       = "stock"
	ORDER       = "order"
	CONTENT     = "content"
)

var ServiceMap = map[string]string{
	"user":       USER,
	"role":       ROLE,
	"permission": PERMISSION,
	"stock":      STOCK,
	"order":      ORDER,
	"content":    CONTENT,
}

type collection struct {
	connMap *hashmap.Map
}

type Registry interface {
	AddAPIEndpoint(key, endpoint string)
	AddAPIConnection(key string, conn *grpc.ClientConn)
	GetAPIConnection(key string) (conn *grpc.ClientConn)
	GetAPIEndpoint(key string) string
	GetUserClient(key string) user.UserServiceClient
	GetRoleClient(key string) user.RoleServiceClient
	GetPermissionClient(key string) user.PermissionServiceClient
	GetStockClient(key string) stock.StockServiceClient
	GetOrderClient(key string) order.OrderServiceClient
	GetContentClient(key string) content.ContentServiceClient
}

// NewRegistry constructs a hashmap for our grpc clients
func NewRegistry() Registry {
	m := hashmap.New()
	return &collection{connMap: m}
}

func (c *collection) AddAPIEndpoint(key, endpoint string) {
	c.connMap.Put(key, endpoint)
}

// AddAPIClient adds a new entry to the hashmap
func (c *collection) AddAPIConnection(key string, conn *grpc.ClientConn) {
	c.connMap.Put(key, conn)
}

// GetAPIClient looks up a client in the hashmap
func (c *collection) GetAPIConnection(key string) (conn *grpc.ClientConn) {
	v, ok := c.connMap.Get(key)
	if !ok {
		panic("could not get grpc client connection")
	}
	conn, _ = v.(*grpc.ClientConn)
	return conn
}

func (c *collection) GetUserClient(key string) user.UserServiceClient {
	return user.NewUserServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetRoleClient(key string) user.RoleServiceClient {
	return user.NewRoleServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetPermissionClient(key string) user.PermissionServiceClient {
	return user.NewPermissionServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetStockClient(key string) stock.StockServiceClient {
	return stock.NewStockServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetOrderClient(key string) order.OrderServiceClient {
	return order.NewOrderServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetContentClient(key string) content.ContentServiceClient {
	return content.NewContentServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetAPIEndpoint(key string) string {
	v, ok := c.connMap.Get(key)
	if !ok {
		panic("could not get api endpoint")
	}
	endpoint, _ := v.(string)
	return endpoint
}
