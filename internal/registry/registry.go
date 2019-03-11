package registry

import (
	pb "github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/emirpasic/gods/maps/hashmap"
)

const (
	USER       = "user"
	ROLE       = "role"
	PERMISSION = "permission"
)

type collection struct {
	clientMap *hashmap.Map
}

type Registry interface {
	AddAPIClient(key string, client interface{})
	GetAPIClient(key string) (interface{}, bool)
	GetUserClient(key string) pb.UserServiceClient
	GetRoleClient(key string) pb.RoleServiceClient
	GetPermissionClient(key string) pb.PermissionServiceClient
}

// NewRegistry constructs a hashmap for our grpc clients
func NewRegistry() Registry {
	m := hashmap.New()
	return &collection{clientMap: m}
}

// AddAPIClient adds a new entry to the hashmap
func (c *collection) AddAPIClient(key string, client interface{}) {
	c.clientMap.Put(key, client)
}

// GetAPIClient looks up a client in the hashmap
func (c *collection) GetAPIClient(key string) (interface{}, bool) {
	return c.clientMap.Get(key)
}

func (c *collection) GetUserClient(key string) pb.UserServiceClient {
	client, ok := c.GetAPIClient(key)
	if !ok {
		panic("could not get user client")
	}
	uc, _ := client.(pb.UserServiceClient)
	return uc
}

func (c *collection) GetRoleClient(key string) pb.RoleServiceClient {
	client, ok := c.GetAPIClient(key)
	if !ok {
		panic("could not get role client")
	}
	rc, _ := client.(pb.RoleServiceClient)
	return rc
}

func (c *collection) GetPermissionClient(key string) pb.PermissionServiceClient {
	client, ok := c.GetAPIClient(key)
	if !ok {
		panic("could not get permission client")
	}
	pc, _ := client.(pb.PermissionServiceClient)
	return pc
}
