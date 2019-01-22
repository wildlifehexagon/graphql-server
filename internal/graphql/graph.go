//go:generate go run ../../scripts/gqlgen.go

package graph

import (
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewGraphQLServer acts as a constructor in dialing GRPC services and returning a defined struct.
func NewGraphQLServer(u string, logger *logrus.Entry) (*resolver.Resolver, error) {
	uconn, err := grpc.Dial(
		u,
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("cannot connect to grpc server for user microservice\n")
		return nil, err
	}
	uc := user.NewUserServiceClient(uconn)

	return &resolver.Resolver{
		UserClient: uc,
		Logger:     logger,
	}, nil
}
