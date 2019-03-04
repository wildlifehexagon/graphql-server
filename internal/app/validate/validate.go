package validate

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v1"
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
