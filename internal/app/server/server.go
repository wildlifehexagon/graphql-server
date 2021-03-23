package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dictyBase/graphql-server/internal/app/middleware"
	"github.com/dictyBase/graphql-server/internal/graphql/dataloader"
	"github.com/dictyBase/graphql-server/internal/graphql/generated"
	"github.com/dictyBase/graphql-server/internal/graphql/resolver"
	"github.com/dictyBase/graphql-server/internal/registry"
	"github.com/dictyBase/graphql-server/internal/repository/redis"
	"github.com/go-chi/cors"
	"google.golang.org/grpc"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// RunGraphQLServer starts the GraphQL backend
func RunGraphQLServer(c *cli.Context) error {
	log := getLogger(c)
	r := chi.NewRouter()
	nr := registry.NewRegistry()
	for k, v := range registry.ServiceMap {
		host := c.String(fmt.Sprintf("%s-grpc-host", k))
		port := c.String(fmt.Sprintf("%s-grpc-port", k))
		// establish grpc connections
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			return cli.NewExitError(
				fmt.Sprintf("cannot connect to grpc microservice %s", err),
				2,
			)
		}
		// add api clients to hashmap
		nr.AddAPIConnection(v, conn)
	}
	endpoints := []string{c.String("publication-api") + "/" + "30048658", c.String("organism-api")}
	// test all api endpoints
	if err := checkEndpoints(endpoints); err != nil {
		return err
	}
	// apis came back ok, add to registry
	nr.AddAPIEndpoint(registry.PUBLICATION, c.String("publication-api"))
	nr.AddAPIEndpoint(registry.ORGANISM, c.String("organism-api"))
	// add redis to registry
	radd := fmt.Sprintf("%s:%s", c.String("redis-master-service-host"), c.String("redis-master-service-port"))
	cache, err := redis.NewCache(radd)
	if err != nil {
		return cli.NewExitError(
			fmt.Sprintf("cannot create redis cache: %v", err),
			2,
		)
	}
	nr.AddRepository("redis", cache)
	s := resolver.NewResolver(nr, log)
	crs := getCORS(c.StringSlice("allowed-origin"))
	r.Use(crs.Handler)
	r.Use(middleware.AuthMiddleWare)
	r.Use(dataloader.DataloaderMiddleware(nr))
	execSchema := generated.NewExecutableSchema(generated.Config{Resolvers: s})
	srv := handler.NewDefaultServer(execSchema)
	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", srv)
	log.Debugf("connect to port 8080 for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", r))
	return nil
}

func checkEndpoints(urls []string) error {
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			return cli.NewExitError(
				fmt.Sprintf("cannot reach api endpoint %s", err),
				2,
			)
		}
		if res.StatusCode != http.StatusOK {
			return cli.NewExitError(
				fmt.Sprintf("did not get ok status from api endpoint, got %v instead", res.StatusCode),
				2,
			)
		}
	}
	return nil
}

func getCORS(origins []string) *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
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
