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

## Development

[glqgen](https://github.com/99designs/gqlgen) relies on GraphQL schema to generate its code. We are storing our schema
files in the [api](./api) folder. There are three files that will be used for each type of schema introduced:

- [query.graphql](./api/query.graphql) - contains all queries
- [mutation.graphql](./api/mutation.graphql) - contains all mutations
- [scalar.graphql](./api/scalar.graphql) - contains any scalars (i.e. timestamp)

Any unique types and inputs for a particular interface (i.e. user) should be placed in their own schema files in the same folder.

In order to generate the actual code, you need to update the [gqlgen.yml](./gqlgen.yml) file. Here are the steps to do so:

1. Add the location for your new schema to the `schema:` field.
2. Change the filename and package for the resolver to something different (i.e. resolvertemp). This has to be done due to limitations with `gqlgen`. From their [Gitter](https://gitter.im/gqlgen/Lobby): _resolver stubbing is currently a "once-only" affair — it wont regenerate if the resolver file exists. This is something we would like to eventually improve, but for now, you will need to manually create new resolvers, by looking at the generated code and adding the missing methods for the resolver interface._
3. Add any (optional) custom models. If using gRPC services, these should be added this way. See the user models for an example.

Now run the generator script with `go generate ./...`. This may take a few minutes, but it will then rebuild the generated files to incorporate your new changes.

Couple of quirks to beware of:

1. The new resolver code does not include the custom package name. So if you are using `generated` as the package name, you are going to have to go in and manually include it like so:

```
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}
```

2. As stated above, you need to go in and compare the diff between the old resolver and the newly generated one, then make sure to add in the changes.
