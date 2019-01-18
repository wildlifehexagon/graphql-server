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
			},
		},
	}
	app.Run(os.Args)
}
