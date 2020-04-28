# Development Guide

- [Development](#development)

  - [Next Steps for Development](#next-steps-for-development)
  - [Error Handling](#error-handling)
  - [Folder Structure](#folder-structure)

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
6. Add any necessary command line flags in [main.go](./cmd/graphql-server/main.go).
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
│   │   └── server
│   │       └── server.go
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
