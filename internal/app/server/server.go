package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dictyBase/go-genproto/dictybaseapis/user"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver"
	"google.golang.org/grpc"

	"github.com/dictyBase/graphql-server/internal/registry"

	"github.com/99designs/gqlgen/handler"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

// RunGraphQLServer starts the GraphQL backend
func RunGraphQLServer(c *cli.Context) error {
	log := getLogger(c)
	// ensure env exists for use in publication resolver
	if len(os.Getenv("PUBLICATION_API_ENDPOINT")) < 1 {
		os.Setenv("PUBLICATION_API_ENDPOINT", c.String("publication-api"))
	}
	// need to think how to optimize this
	// use NewRegistry to create connections?
	uconn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", c.String("user-grpc-host"), c.String("user-grpc-port")),
		grpc.WithInsecure(),
	)
	if err != nil {
		return cli.NewExitError(
			fmt.Sprintf("cannot connect to user grpc microservice %s", err.Error()),
			2,
		)
	}
	uc := user.NewUserServiceClient(uconn)
	rconn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", c.String("role-grpc-host"), c.String("role-grpc-port")),
		grpc.WithInsecure(),
	)
	if err != nil {
		return cli.NewExitError(
			fmt.Sprintf("cannot connect to role grpc microservice %s", err.Error()),
			2,
		)
	}
	rc := user.NewRoleServiceClient(rconn)
	pconn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", c.String("permission-grpc-host"), c.String("permission-grpc-port")),
		grpc.WithInsecure(),
	)
	if err != nil {
		return cli.NewExitError(
			fmt.Sprintf("cannot connect to permission grpc microservice %s", err.Error()),
			2,
		)
	}
	pc := user.NewPermissionServiceClient(pconn)

	nr := registry.NewRegistry()
	nr.AddAPIClient("user", uc)
	nr.AddAPIClient("role", rc)
	nr.AddAPIClient("permission", pc)

	s := resolver.NewResolver(nr, log)

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: s})))

	log.Debug("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}

func getLogger(c *cli.Context) *logrus.Entry {
	log := logrus.New()
	log.Out = os.Stderr
	switch c.GlobalString("log-format") {
	case "text":
		log.Formatter = &logrus.TextFormatter{
			TimestampFormat: "02/Jan/2006:15:04:05",
		}
	case "json":
		log.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "02/Jan/2006:15:04:05",
		}
	}
	l := c.GlobalString("log-level")
	switch l {
	case "debug":
		log.Level = logrus.DebugLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	case "panic":
		log.Level = logrus.PanicLevel
	}
	return logrus.NewEntry(log)
}
