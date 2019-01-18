//go:generate go run ../../scripts/gqlgen.go
package graph

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

// Client groups together the gRPC service clients.
type Client struct {
	userClient user.UserServiceClient
	logger     *logrus.Entry
}

// NewGraphQLServer acts as a constructor in dialing GRPC services and returning a defined struct.
func NewGraphQLServer(u string, logger *logrus.Entry) (*Client, error) {
	// connect to user service
	uconn, err := grpc.Dial(
		u,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, cli.NewExitError(
			fmt.Sprintf("cannot connect to grpc server for user microservice %s", err),
			2,
		)
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
