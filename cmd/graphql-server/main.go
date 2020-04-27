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
			Flags:  getServerFlags(),
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error in running command %s", err)
	}
}

func getServerFlags() []cli.Flag {
	var f []cli.Flag
	f = append(f, userFlags()...)
	f = append(f, redisFlags()...)
	f = append(f, dscFlags()...)
	f = append(f, contentFlags()...)
	f = append(f, nonGRPCFlags()...)
	f = append(f, allowedOriginFlags()...)
	return append(f, authFlags()...)
}

func userFlags() []cli.Flag {
	return []cli.Flag{
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
	}
}

func redisFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "redis-master-service-host",
			EnvVar: "REDIS_MASTER_SERVICE_HOST",
			Usage:  "redis master grpc host",
		},
		cli.StringFlag{
			Name:   "redis-master-service-port",
			EnvVar: "REDIS_MASTER_SERVICE_PORT",
			Usage:  "redis master grpc port",
		},
		cli.IntFlag{
			Name:  "cache-expiration-days",
			Usage: "number of days to store redis cache",
			Value: 7,
		},
	}
}

func authFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "auth-grpc-host",
			EnvVar: "AUTH_API_SERVICE_HOST",
			Usage:  "auth grpc host",
		},
		cli.StringFlag{
			Name:   "auth-grpc-port",
			EnvVar: "AUTH_API_SERVICE_PORT",
			Usage:  "auth grpc port",
		},
		cli.StringFlag{
			Name:   "identity-grpc-host",
			EnvVar: "IDENTITY_API_SERVICE_HOST",
			Usage:  "identity grpc host",
		},
		cli.StringFlag{
			Name:   "identity-grpc-port",
			EnvVar: "IDENTITY_API_SERVICE_PORT",
			Usage:  "identity grpc port",
		},
	}
}

func dscFlags() []cli.Flag {
	return []cli.Flag{
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
	}
}

func contentFlags() []cli.Flag {
	return []cli.Flag{
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
	}
}

func nonGRPCFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "publication-api, pub",
			EnvVar: "PUBLICATION_API_ENDPOINT",
			Usage:  "publication api endpoint",
		},
	}
}

/**
  A list of allowed origins is necessary since server has set 
  allow-credentials to true.
  See https://github.com/rs/cors/issues/55
*/
func allowedOriginFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringSliceFlag{
			Name:   "allowed-origin",
			Usage: "allowed origins for CORS",
			Value: &cli.StringSlice{},
		},
	}
}