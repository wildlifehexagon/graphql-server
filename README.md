# graphql-server

[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](LICENSE)  
![GitHub action](https://github.com/dictyBase/graphql-server/workflows/Build/badge.svg)
[![codecov](https://codecov.io/gh/dictyBase/modware-annotation/branch/develop/graph/badge.svg)](https://codecov.io/gh/dictyBase/modware-annotation)  
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

[Full development guide](./docs/development.md) available in the `docs` folder.

## Active Developers

<a href="https://sourcerer.io/cybersiddhu"><img src="https://sourcerer.io/assets/avatar/cybersiddhu" height="80px" alt="Sourcerer"></a>
<a href="https://sourcerer.io/wildlifehexagon"><img src="https://sourcerer.io/assets/avatar/wildlifehexagon" height="80px" alt="Sourcerer"></a>
