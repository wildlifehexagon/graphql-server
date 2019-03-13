package registry

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/emirpasic/gods/maps/hashmap"
	"google.golang.org/grpc"
)

const (
	USER        = "user"
	ROLE        = "role"
	PERMISSION  = "permission"
	PUBLICATION = "publication"
)

var ServiceMap = map[string]string{
	"user":       USER,
	"role":       ROLE,
	"permission": PERMISSION,
}

type collection struct {
	connMap *hashmap.Map
}

type Registry interface {
	AddAPIEndpoint(key, endpoint string)
	AddAPIConnection(key string, conn *grpc.ClientConn)
	GetAPIConnection(key string) (conn *grpc.ClientConn)
	GetAPIEndpoint(key string) string
	GetUserClient(key string) pb.UserServiceClient
	GetRoleClient(key string) pb.RoleServiceClient
	GetPermissionClient(key string) pb.PermissionServiceClient
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

func (c *collection) GetUserClient(key string) pb.UserServiceClient {
	return user.NewUserServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetRoleClient(key string) pb.RoleServiceClient {
	return user.NewRoleServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetPermissionClient(key string) pb.PermissionServiceClient {
	return user.NewPermissionServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetAPIEndpoint(key string) string {
	v, ok := c.connMap.Get(key)
	if !ok {
		panic("could not get api endpoint")
	}
	endpoint, _ := v.(string)
	return endpoint
}
