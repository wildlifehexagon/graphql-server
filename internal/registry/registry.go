package registry

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/annotation"
	"github.com/dictyBase/go-genproto/dictybaseapis/auth"
	"github.com/dictyBase/go-genproto/dictybaseapis/content"
	"github.com/dictyBase/go-genproto/dictybaseapis/identity"
	"github.com/dictyBase/go-genproto/dictybaseapis/order"
	"github.com/dictyBase/go-genproto/dictybaseapis/stock"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
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

type Registry struct {
	connMap *hashmap.Map
}

// NewRegistry constructs a hashmap for our grpc clients
func NewRegistry() Registry {
	m := hashmap.New()
	return Registry{connMap: m}
}

func (r *Registry) AddAPIEndpoint(key, endpoint string) {
	r.connMap.Put(key, endpoint)
}

// AddAPIClient adds a new entry to the hashmap
func (r *Registry) AddAPIConnection(key string, conn *grpc.ClientConn) {
	r.connMap.Put(key, conn)
}

// GetAPIClient looks up a client in the hashmap
func (r *Registry) GetAPIConnection(key string) (conn *grpc.ClientConn) {
	v, ok := r.connMap.Get(key)
	if !ok {
		panic("could not get grpc client connection")
	}
	conn, _ = v.(*grpc.ClientConn)
	return conn
}

func (r *Registry) GetUserClient(key string) user.UserServiceClient {
	return user.NewUserServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetRoleClient(key string) user.RoleServiceClient {
	return user.NewRoleServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetPermissionClient(key string) user.PermissionServiceClient {
	return user.NewPermissionServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetStockClient(key string) stock.StockServiceClient {
	return stock.NewStockServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetOrderClient(key string) order.OrderServiceClient {
	return order.NewOrderServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetContentClient(key string) content.ContentServiceClient {
	return content.NewContentServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetAnnotationClient(key string) annotation.TaggedAnnotationServiceClient {
	return annotation.NewTaggedAnnotationServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetAuthClient(key string) auth.AuthServiceClient {
	return auth.NewAuthServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetIdentityClient(key string) identity.IdentityServiceClient {
	return identity.NewIdentityServiceClient(r.GetAPIConnection(key))
}

func (r *Registry) GetAPIEndpoint(key string) string {
	v, ok := r.connMap.Get(key)
	if !ok {
		panic("could not get api endpoint")
	}
	endpoint, _ := v.(string)
	return endpoint
}
