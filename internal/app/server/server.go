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
	// establish grpc connections
	uconn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.String("user-grpc-host"), c.String("user-grpc-port")), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to grpc user microservice for %s", err)
	}
	rconn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.String("role-grpc-host"), c.String("role-grpc-port")), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to grpc role microservice for %s", err)
	}
	pconn, err := grpc.Dial(fmt.Sprintf("%s:%s", c.String("permission-grpc-host"), c.String("permission-grpc-port")), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("cannot connect to grpc permission microservice for %s", err)
	}
	// generate new (empty) hashmap
	nr := registry.NewRegistry()
	// add api clients to hashmap
	nr.AddAPIClient(registry.USER, user.NewUserServiceClient(uconn))
	nr.AddAPIClient(registry.ROLE, user.NewRoleServiceClient(rconn))
	nr.AddAPIClient(registry.PERMISSION, user.NewPermissionServiceClient(pconn))

	s := resolver.NewResolver(nr, log)

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: s})))
	log.Debugf("connect to http://localhost:8080/ for GraphQL playground")
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
