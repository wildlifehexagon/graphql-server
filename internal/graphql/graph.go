//go:generate go run ../../scripts/gqlgen.go

package graph

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	resolver "github.com/dictyBase/graphql-server/internal/graphql/resolver"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewGraphQLServer acts as a constructor in dialing GRPC services and returning a defined struct.
func NewGraphQLServer(u, r, p string, logger *logrus.Entry) (*resolver.Resolver, error) {
	uconn, err := grpc.Dial(
		u,
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("cannot connect to grpc server for user microservice\n")
		return nil, err
	}
	uc := user.NewUserServiceClient(uconn)
	rconn, err := grpc.Dial(
		r,
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("cannot connect to grpc server for role microservice\n")
		return nil, err
	}
	rc := user.NewRoleServiceClient(rconn)
	pconn, err := grpc.Dial(
		p,
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
