//go:generate go run ../../scripts/gqlgen.go
package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Client groups together the gRPC service clients.
type Client struct {
	userClient user.UserServiceClient
	logger     *logrus.Entry
}

// NewGraphQLServer acts as a constructor in dialing GRPC services and returning a defined struct.
func NewGraphQLServer(u string, logger *logrus.Entry) (*Client, error) {
	uconn, err := grpc.Dial(
		u,
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.Fatalf("cannot connect to grpc server for user microservice\n")
		return nil, err
	}
	uc := user.NewUserServiceClient(uconn)

	return &Client{
		uc,
		logger,
	}, nil
}

func (c *Client) Mutation() MutationResolver {
	return &mutationResolver{
		client: c,
	}
}

func (c *Client) Query() QueryResolver {
	return &queryResolver{
		client: c,
	}
}

func (c *Client) ToExecutableSchema() graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: c,
	})
}
