package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	graph "github.com/dictyBase/graphql-server/internal/graphql"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const defaultPort = "8080"

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

// RunGraphQLServer starts the GraphQL backend
func RunGraphQLServer(c *cli.Context) error {
	l := getLogger(c)

	u := fmt.Sprintf("%s:%s", c.String("user-grpc-host"), c.String("user-grpc-port"))
	g, err := graph.NewGraphQLServer(u, l)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(g.ToExecutableSchema()))

	l.Info("connect to http://localhost:8080/ for GraphQL playground")
	l.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
