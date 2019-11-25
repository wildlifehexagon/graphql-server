package main

import (
	"log"
	"os"

	"github.com/dictyBase/graphql-server/internal/app/server"
	"github.com/dictyBase/graphql-server/internal/app/validate"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "graphql-server"
	app.Usage = "cli for graphql-server backend"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log-format",
			Usage: "format of the logging out, either of json or text.",
			Value: "json",
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "log level for the application",
			Value: "error",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "start-server",
			Usage:  "starts the graphql-server backend",
			Action: server.RunGraphQLServer,
			Before: validate.ValidateServerArgs,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:   "user-grpc-host",
					EnvVar: "USER_API_SERVICE_HOST",
					Usage:  "user grpc host",
				},
				cli.StringFlag{
					Name:   "user-grpc-port",
					EnvVar: "USER_API_SERVICE_PORT",
					Usage:  "user grpc port",
				},
				cli.StringFlag{
					Name:   "role-grpc-host",
					EnvVar: "ROLE_API_SERVICE_HOST",
					Usage:  "role grpc host",
				},
				cli.StringFlag{
					Name:   "role-grpc-port",
					EnvVar: "ROLE_API_SERVICE_PORT",
					Usage:  "role grpc port",
				},
				cli.StringFlag{
					Name:   "permission-grpc-host",
					EnvVar: "PERMISSION_API_SERVICE_HOST",
					Usage:  "permission grpc host",
				},
				cli.StringFlag{
					Name:   "permission-grpc-port",
					EnvVar: "PERMISSION_API_SERVICE_PORT",
					Usage:  "permission grpc port",
				},
				cli.StringFlag{
					Name:   "content-grpc-host",
					EnvVar: "CONTENT_API_SERVICE_HOST",
					Usage:  "content grpc host",
				},
				cli.StringFlag{
					Name:   "content-grpc-port",
					EnvVar: "CONTENT_API_SERVICE_PORT",
					Usage:  "content grpc port",
				},
				cli.StringFlag{
					Name:   "publication-api, pub",
					EnvVar: "PUBLICATION_API_ENDPOINT",
					Usage:  "publication api endpoint",
				},
				cli.StringFlag{
					Name:   "stock-grpc-host",
					EnvVar: "STOCK_API_SERVICE_HOST",
					Usage:  "stock grpc host",
				},
				cli.StringFlag{
					Name:   "stock-grpc-port",
					EnvVar: "STOCK_API_SERVICE_PORT",
					Usage:  "stock grpc port",
				},
				cli.StringFlag{
					Name:   "order-grpc-host",
					EnvVar: "ORDER_API_SERVICE_HOST",
					Usage:  "order grpc host",
				},
				cli.StringFlag{
					Name:   "order-grpc-port",
					EnvVar: "ORDER_API_SERVICE_PORT",
					Usage:  "order grpc port",
				},
				cli.StringFlag{
					Name:   "annotation-grpc-host",
					EnvVar: "ANNOTATION_API_SERVICE_HOST",
					Usage:  "annotation grpc host",
				},
				cli.StringFlag{
					Name:   "annotation-grpc-port",
					EnvVar: "ANNOTATION_API_SERVICE_PORT",
					Usage:  "annotation grpc port",
				},
				cli.StringFlag{
					Name:   "redis-master-service-host",
					EnvVar: "REDIS_MASTER_SERVICE_HOST",
					Usage:  "redis master service host",
				},
				cli.StringFlag{
					Name:   "redis-master-service-port",
					EnvVar: "REDIS_MASTER_SERVICE_PORT",
					Usage:  "redis master service port",
				},
				cli.IntFlag{
					Name:  "cache-expiration-days",
					Usage: "number of days to store redis cache",
					Value: 7,
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
