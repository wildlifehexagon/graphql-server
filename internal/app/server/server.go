package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dictyBase/graphql-server/internal/registry"

	"github.com/99designs/gqlgen/handler"
	graph "github.com/dictyBase/graphql-server/internal/graphql"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

// RunGraphQLServer starts the GraphQL backend
func RunGraphQLServer(c *cli.Context) error {
	log := getLogger(c)

	nr := registry.NewRegistry()
	nr.AddAPIClient("user", fmt.Sprintf("%s:%s", c.String("user-grpc-host"), c.String("user-grpc-port")))
	nr.AddAPIClient("role", fmt.Sprintf("%s:%s", c.String("role-grpc-host"), c.String("role-grpc-port")))
	nr.AddAPIClient("permission", fmt.Sprintf("%s:%s", c.String("permission-grpc-host"), c.String("permission-grpc-port")))

	s, err := graph.NewGraphQLServer(nr, log)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: s})))

	log.Info("connect to http://localhost:8080/ for GraphQL playground")
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
