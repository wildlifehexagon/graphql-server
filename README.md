# graphql-server

dictyBase GraphQL server. This uses [glqgen](https://github.com/99designs/gqlgen) to generate code to match our gRPC models.

## Usage

```
NAME:
   graphql-server - cli for graphql-server backend

USAGE:
   graphql-server [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     start-server  starts the graphql-server backend
     help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-format value  format of the logging out, either of json or text. (default: "json")
   --log-level value   log level for the application (default: "error")
   --help, -h          show help
   --version, -v       print the version
```

## Subcommand

```
NAME:
   graphql-server start-server - starts the graphql-server backend

USAGE:
   graphql-server start-server [command options] [arguments...]

OPTIONS:
   --user-grpc-host value, --uh value        user grpc host [$USER_API_SERVICE_HOST]
   --user-grpc-port value, --up value        user grpc port [$USER_API_SERVICE_PORT]
   --role-grpc-host value, --rh value        role grpc host [$ROLE_API_SERVICE_HOST]
   --role-grpc-port value, --rp value        role grpc port [$ROLE_API_SERVICE_PORT]
   --permission-grpc-host value, --ph value  permission grpc host [$PERMISSION_API_SERVICE_HOST]
   --permission-grpc-port value, --pp value  permission grpc port [$PERMISSION_API_SERVICE_PORT]
```
