package main

import (
	"os"

	"github.com/dictyBase/graphql-server/internal/app/server"
	"github.com/dictyBase/graphql-server/internal/app/validate"
	cli "gopkg.in/urfave/cli.v1"
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
					Name:   "user-grpc-host, uh",
					EnvVar: "USER_API_SERVICE_HOST",
					Usage:  "user grpc host",
				},
				cli.StringFlag{
					Name:   "user-grpc-port, up",
					EnvVar: "USER_API_SERVICE_PORT",
					Usage:  "user grpc port",
				},
				cli.StringFlag{
					Name:   "role-grpc-host, rh",
					EnvVar: "ROLE_API_SERVICE_HOST",
					Usage:  "role grpc host",
				},
				cli.StringFlag{
					Name:   "role-grpc-port, rp",
					EnvVar: "ROLE_API_SERVICE_PORT",
					Usage:  "role grpc port",
				},
				cli.StringFlag{
					Name:   "permission-grpc-host, ph",
					EnvVar: "PERMISSION_API_SERVICE_HOST",
					Usage:  "permission grpc host",
				},
				cli.StringFlag{
					Name:   "permission-grpc-port, pp",
					EnvVar: "PERMISSION_API_SERVICE_PORT",
					Usage:  "permission grpc port",
				},
				cli.StringFlag{
					Name:   "content-grpc-host, ch",
					EnvVar: "CONTENT_API_SERVICE_HOST",
					Usage:  "content grpc host",
				},
				cli.StringFlag{
					Name:   "content-grpc-port, cp",
					EnvVar: "CONTENT_API_SERVICE_PORT",
					Usage:  "content grpc port",
				},
				cli.StringFlag{
					Name:   "publication-api, pub",
					EnvVar: "PUBLICATION_API_ENDPOINT",
					Usage:  "publication api endpoint",
					Value:  "https://betafunc.dictybase.org/publications",
				},
				cli.StringFlag{
					Name:   "stock-grpc-host, sh",
					EnvVar: "STOCK_API_SERVICE_HOST",
					Usage:  "stock grpc host",
				},
				cli.StringFlag{
					Name:   "stock-grpc-port, sp",
					EnvVar: "STOCK_API_SERVICE_PORT",
					Usage:  "stock grpc port",
				},
				cli.StringFlag{
					Name:   "order-grpc-host, oh",
					EnvVar: "ORDER_API_SERVICE_HOST",
					Usage:  "order grpc host",
				},
				cli.StringFlag{
					Name:   "order-grpc-port, op",
					EnvVar: "ORDER_API_SERVICE_PORT",
					Usage:  "order grpc port",
				},
			},
		},
	}
	app.Run(os.Args)
}
