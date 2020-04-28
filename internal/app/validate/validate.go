package validate

import (
	"fmt"

	"github.com/urfave/cli"
)

// ValidateServerArgs validates that the necessary flags are not missing
func ValidateServerArgs(c *cli.Context) error {
	for _, p := range []string{
		"user-grpc-host",
		"user-grpc-port",
		"role-grpc-host",
		"role-grpc-port",
		"permission-grpc-host",
		"permission-grpc-port",
		"stock-grpc-host",
		"stock-grpc-port",
		"order-grpc-host",
		"order-grpc-port",
		"content-grpc-host",
		"content-grpc-port",
		"annotation-grpc-host",
		"annotation-grpc-port",
		"auth-grpc-host",
		"auth-grpc-port",
		"identity-grpc-host",
		"identity-grpc-port",
	} {
		if len(c.String(p)) == 0 {
			return cli.NewExitError(
				fmt.Sprintf("argument %s is missing", p),
				2,
			)
		}
	}
	return nil
}
