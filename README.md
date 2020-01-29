# graphql-server

[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](LICENSE)  
![GitHub action](https://github.com/dictyBase/graphql-server/workflows/Test%20coverage/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/dictyBase/graphql-server)](https://goreportcard.com/report/github.com/dictyBase/graphql-server)
[![Technical debt](https://badgen.net/codeclimate/tech-debt/dictyBase/graphql-server)](https://codeclimate.com/github/dictyBase/graphql-server/trends/technical_debt)
[![Issues](https://badgen.net/codeclimate/issues/dictyBase/graphql-server)](https://codeclimate.com/github/dictyBase/graphql-server/issues)
[![Maintainability](https://api.codeclimate.com/v1/badges/21ed283a6186cfa3d003/maintainability)](https://codeclimate.com/github/dictyBase/graphql-server/maintainability)
[![Dependabot Status](https://api.dependabot.com/badges/status?host=github&repo=dictyBase/graphql-server)](https://dependabot.com)  
![Issues](https://badgen.net/github/issues/dictyBase/graphql-server)
![Open Issues](https://badgen.net/github/open-issues/dictyBase/graphql-server)
![Closed Issues](https://badgen.net/github/closed-issues/dictyBase/graphql-server)  
![Total PRS](https://badgen.net/github/prs/dictyBase/graphql-server)
![Open PRS](https://badgen.net/github/open-prs/dictyBase/graphql-server)
![Closed PRS](https://badgen.net/github/closed-prs/dictyBase/graphql-server)
![Merged PRS](https://badgen.net/github/merged-prs/dictyBase/graphql-server)  
![Commits](https://badgen.net/github/commits/dictyBase/graphql-server/develop)
![Last commit](https://badgen.net/github/last-commit/dictyBase/graphql-server/develop)
![Branches](https://badgen.net/github/branches/dictyBase/graphql-server)
![Tags](https://badgen.net/github/tags/dictyBase/graphql-server/?color=cyan)  
![GitHub repo size](https://img.shields.io/github/repo-size/dictyBase/graphql-server?style=plastic)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/dictyBase/graphql-server?style=plastic)
[![Lines of Code](https://badgen.net/codeclimate/loc/dictyBase/graphql-server)](https://codeclimate.com/github/dictyBase/graphql-server/code)  
[![Funding](https://badgen.net/badge/NIGMS/Rex%20L%20Chisholm,dictyBase/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=9476993)
[![Funding](https://badgen.net/badge/NIGMS/Rex%20L%20Chisholm,DSC/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=9438930)

dictyBase GraphQL server. This uses [glqgen](https://github.com/99designs/gqlgen) to generate code to match our gRPC models.

- [Usage](#usage)
- [Subcommand](#subcommand)
- [Development](#development)
  - [Next Steps for Development](#next-steps-for-development)
  - [Error Handling](#error-handling)
  - [Folder Structure](#folder-structure)

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
   --publication-api value, --pub value      publication api endpoint (default: "https://betafunc.dictybase.org/publications") [$PUBLICATION_API_ENDPOINT]
   --stock-grpc-host value, --sh value       stock grpc host [$STOCK_API_SERVICE_HOST]
   --stock-grpc-port value, --sp value       stock grpc port [$STOCK_API_SERVICE_PORT]
   --order-grpc-host value, --oh value       order grpc host [$ORDER_API_SERVICE_HOST]
   --order-grpc-port value, --op value       order grpc port [$ORDER_API_SERVICE_PORT]
```

## Development

[gqlgen](https://github.com/99designs/gqlgen) relies on GraphQL schema to generate its code. We are storing our schema
files in the [dictyBase/graphql-schema](https://github.com/dictyBase/graphql-schema) repository. There are three files that will be used for each type of schema introduced:

- [query.graphql](https://github.com/dictyBase/graphql-schema/blob/master/query.graphql) - contains all queries
- [mutation.graphql](https://github.com/dictyBase/graphql-schema/blob/master/mutation.graphql) - contains all mutations
- [scalar.graphql](https://github.com/dictyBase/graphql-schema/blob/master/scalar.graphql) - contains any scalars (i.e. timestamp)

Any unique types and inputs for a particular category (i.e. user) should be placed in their own schema files in the same folder.

In order to generate the actual code, you need to update the [gqlgen.yml](./gqlgen.yml) file. Here are the steps to do so:

1. Add the location for your new schema to the `schema:` field.
2. Add any (optional) custom models. If using gRPC services, these should be added this way. See the user models for an example.

Now run the generator script with `go generate ./...`. This may take a few minutes, but it will then rebuild the generated files to incorporate your new changes.

Couple of quirks to beware of:

1. The new resolver code does not include the custom package name. So if you are using `generated` as the package name, you are going to have to go in and manually include it like so:

```
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}
```

2. You need to go in and compare the diff between the old resolver and the newly generated one, then make sure to add in the changes.

### Next Steps for Development

After adding your new schema and running the generator script, you will need to do a bit of additional refactoring.

1. Add appropriate constant to [registry.go](./internal/registry/registry.go) for that client.
2. If adding a gRPC client, you also need to add an entry to the `ServiceMap` in the same file. If adding an API endpoint, you need to update [RunGraphQLServer](./internal/app/server/server.go) to manually add this to the registry. See how this was done with Publication for an example. Also be sure to test the endpoint during server initialization.
3. Add a new method to the [registry](./internal/registry/registry.go) package for that client (a la `GetUserClient`).
4. Next, look at the newly generated resolvers from your schema. Move any _shared_ resolvers into the root [resolver](./internal/graphql/resolver) folder.
   - Put the main generated function (i.e. `Permission()`) in the [resolver.go](./internal/graphql/resolver/resolver.go) file and update it to match the format of the other functions.
   - Create new files as necessary for each client, each one containing their query and mutation methods. Look at [permission.go](./internal/graphql/resolver/permission.go) for an example. Make sure to update the receivers for each method. For queries we are using `(q *QueryResolver)` and for mutations `(m *MutationResolver)`.
5. Now add any _unshared_ resolvers into a separate folder inside `resolver`. These resolvers are generally tied to the individual fields for that model, and they are unique to that particular client. You can look at the [user](./internal/graphql/resolver/user) folder for examples. Also update the package name if necessary.
6. Add any necessary command line flags in [main.go](./cmd/graphql-server/main.go) and [validation](./internal/app/validate/validate.go).
7. Fill out your resolver stubs and then test it out in the playground!

### Error Handling

To improve error handling in our front end web applications, we are adding custom error messages when applicable on the server side. The `errorutils` function in the package of the same name is used to add new errors with custom extensions. See any of the shared resolvers for an example in how to use this.

### Folder Structure

```
.
├── README.md
├── build
│   └── Dockerfile
├── cmd
│   └── graphql-server
│       └── main.go
├── deployments
│   └── charts
│       └── graphql-server
│           ├── Chart.yaml
│           ├── templates
│           │   ├── NOTES.txt
│           │   ├── _helpers.tpl
│           │   ├── deployment.yaml
│           │   └── service.yaml
│           └── values.yaml
├── go.mod
├── go.sum
├── gqlgen.yml
├── graphql-server
├── internal
│   ├── app
│   │   ├── server
│   │   │   └── server.go
│   │   └── validate
│   │       └── validate.go
│   ├── graphql
│   │   ├── generated
│   │   │   └── generated.go
│   │   ├── models
│   │   │   ├── models_gen.go
│   │   │   └── timestamp.go
│   │   └── resolver
│   │       ├── permission.go
│   │       ├── resolver.go
│   │       ├── role.go
│   │       ├── user
│   │       │   ├── permission.go
│   │       │   ├── role.go
│   │       │   └── user.go
│   │       └── user.go
│   └── registry
│       └── registry.go
└── scripts
    └── gqlgen.go
```

## Active Developers

<a href="https://sourcerer.io/cybersiddhu"><img src="https://sourcerer.io/assets/avatar/cybersiddhu" height="80px" alt="Sourcerer"></a>
<a href="https://sourcerer.io/wildlifehexagon"><img src="https://sourcerer.io/assets/avatar/wildlifehexagon" height="80px" alt="Sourcerer"></a>
