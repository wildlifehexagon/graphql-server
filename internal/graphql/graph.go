//go:generate go run ../../scripts/gqlgen.go

package graph

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	resolver "github.com/dictyBase/graphql-server/internal/graphql/resolver"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewGraphQLServer acts as a constructor in dialing GRPC services and returning a defined struct.
func NewGraphQLServer(nr registry.Registry, logger *logrus.Entry) (*resolver.Resolver, error) {
	// needs to be optimized
	// code complexity will be too high soon
	u, _ := nr.GetAPIClient("user")
	uconn, err := grpc.Dial(
		u.(string),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("cannot connect to grpc server for user microservice\n")
		return nil, err
	}
	uc := user.NewUserServiceClient(uconn)
	r, _ := nr.GetAPIClient("role")
	rconn, err := grpc.Dial(
		r.(string),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("cannot connect to grpc server for role microservice\n")
		return nil, err
	}
	rc := user.NewRoleServiceClient(rconn)
	p, _ := nr.GetAPIClient("permission")
	pconn, err := grpc.Dial(
		p.(string),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("cannot connect to grpc server for permission microservice\n")
		return nil, err
	}
	pc := user.NewPermissionServiceClient(pconn)

	return &resolver.Resolver{
		UserClient:       uc,
		RoleClient:       rc,
		PermissionClient: pc,
		Logger:           logger,
	}, nil
}
