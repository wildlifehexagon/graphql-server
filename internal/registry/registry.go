package registry

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/go-genproto/dictybaseapis/identity"
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/repository"
	"github.com/emirpasic/gods/maps/hashmap"
	"google.golang.org/grpc"
)

// constants used with modware-annotation
const (
	PhenoOntology       = "Dicty Phenotypes"
	EnvOntology         = "Dicty Environment"
	AssayOntology       = "Dictyostelium Assay"
	MutagenesisOntology = "Dd Mutagenesis Method"
	DictyAnnoOntology   = "dicty_annotation"
	StrainCharOnto      = "strain_characteristics"
	StrainInvOnto       = "strain_inventory"
	PlasmidInvOnto      = "plasmid_inventory"
	StrainInvTag        = "strain_inventory"
	PlasmidInvTag       = "plasmid inventory"
	InvLocationTag      = "location"
	LiteratureTag       = "pubmed id"
	NoteTag             = "public note"
	SysnameTag          = "systematic name"
	MutmethodTag        = "mutagenesis method"
	MuttypeTag          = "mutant type"
	GenoTag             = "genotype"
	SynTag              = "synonym"
	EmptyValue          = "novalue"
)

const (
	USER        = "user"
	ROLE        = "role"
	PERMISSION  = "permission"
	PUBLICATION = "publication"
	STOCK       = "stock"
	ORDER       = "order"
	CONTENT     = "content"
	ANNOTATION  = "annotation"
	AUTH        = "auth"
	IDENTITY    = "identity"
)

var ServiceMap = map[string]string{
	"user":       USER,
	"role":       ROLE,
	"permission": PERMISSION,
	"stock":      STOCK,
	"order":      ORDER,
	"content":    CONTENT,
	"annotation": ANNOTATION,
	"auth":       AUTH,
	"identity":   IDENTITY,
}

type collection struct {
	connMap *hashmap.Map
}

type Registry interface {
	AddAPIEndpoint(key, endpoint string)
	AddAPIConnection(key string, conn *grpc.ClientConn)
	AddRepository(key string, st repository.Repository)
	GetAPIConnection(key string) (conn *grpc.ClientConn)
	GetAPIEndpoint(key string) string
	GetUserClient(key string) user.UserServiceClient
	GetRoleClient(key string) user.RoleServiceClient
	GetPermissionClient(key string) user.PermissionServiceClient
	GetStockClient(key string) stock.StockServiceClient
	GetOrderClient(key string) order.OrderServiceClient
	GetContentClient(key string) content.ContentServiceClient
	GetAnnotationClient(key string) annotation.TaggedAnnotationServiceClient
	GetAuthClient(key string) auth.AuthServiceClient
	GetIdentityClient(key string) identity.IdentityServiceClient
	GetRedisRepository(key string) repository.Repository
}

// NewRegistry constructs a hashmap for our grpc clients
func NewRegistry() Registry {
	m := hashmap.New()
	return &collection{connMap: m}
}

// AddAPIEndpoint adds a new REST endpoint to the hashmap
func (c *collection) AddAPIEndpoint(key, endpoint string) {
	c.connMap.Put(key, endpoint)
}

// AddAPIClient adds a new gRPC client to the hashmap
func (c *collection) AddAPIConnection(key string, conn *grpc.ClientConn) {
	c.connMap.Put(key, conn)
}

// AddRepository adds a new repository client to the hashmap
func (c *collection) AddRepository(key string, st repository.Repository) {
	c.connMap.Put(key, st)
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

func (c *collection) GetAnnotationClient(key string) annotation.TaggedAnnotationServiceClient {
	return annotation.NewTaggedAnnotationServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetAuthClient(key string) auth.AuthServiceClient {
	return auth.NewAuthServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetIdentityClient(key string) identity.IdentityServiceClient {
	return identity.NewIdentityServiceClient(c.GetAPIConnection(key))
}

func (c *collection) GetAPIEndpoint(key string) string {
	v, _ := c.connMap.Get(key)
	endpoint, _ := v.(string)
	return endpoint
}

func (c *collection) GetRedisRepository(key string) repository.Repository {
	v, _ := c.connMap.Get(key)
	st, _ := v.(repository.Repository)
	return st
}
