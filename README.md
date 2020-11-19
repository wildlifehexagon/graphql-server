# graphql-server
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

[![License](https://img.shields.io/badge/License-BSD%202--Clause-blue.svg)](LICENSE)  
![GitHub action](https://github.com/dictyBase/graphql-server/workflows/Build/badge.svg)
[![codecov](https://codecov.io/gh/dictyBase/graphql-server/branch/develop/graph/badge.svg)](https://codecov.io/gh/dictyBase/graphql-server)  
[![Maintainability](https://api.codeclimate.com/v1/badges/21ed283a6186cfa3d003/maintainability)](https://codeclimate.com/github/dictyBase/graphql-server/maintainability)  
![Last commit](https://badgen.net/github/last-commit/dictyBase/graphql-server/develop)
[![Funding](https://badgen.net/badge/Funding/Rex%20L%20Chisholm,dictyBase,DCR/yellow?list=|)](https://projectreporter.nih.gov/project_info_description.cfm?aid=10024726&icde=0)

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

In order to use the GraphQL Playground, you will need to add the following custom HTTP header:

```json
{
  "X-GraphQL-Method": "Query"
}
```

# Misc badges
![Issues](https://badgen.net/github/issues/dictyBase/graphql-server)
![Open Issues](https://badgen.net/github/open-issues/dictyBase/graphql-server)
![Closed Issues](https://badgen.net/github/closed-issues/dictyBase/graphql-server)  
![Total PRS](https://badgen.net/github/prs/dictyBase/graphql-server)
![Open PRS](https://badgen.net/github/open-prs/dictyBase/graphql-server)
![Closed PRS](https://badgen.net/github/closed-prs/dictyBase/graphql-server)
![Merged PRS](https://badgen.net/github/merged-prs/dictyBase/graphql-server)  
![Commits](https://badgen.net/github/commits/dictyBase/graphql-server/develop)
![Branches](https://badgen.net/github/branches/dictyBase/graphql-server)
![Tags](https://badgen.net/github/tags/dictyBase/graphql-server/?color=cyan)  
![GitHub repo size](https://img.shields.io/github/repo-size/dictyBase/graphql-server?style=plastic)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/dictyBase/graphql-server?style=plastic)
[![Lines of Code](https://badgen.net/codeclimate/loc/dictyBase/graphql-server)](https://codeclimate.com/github/dictyBase/graphql-server/code)  

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://cybersiddhu.github.com/"><img src="https://avatars3.githubusercontent.com/u/48740?v=4" width="100px;" alt=""/><br /><sub><b>Siddhartha Basu</b></sub></a><br /><a href="https://github.com/dictyBase/graphql-server/issues?q=author%3Acybersiddhu" title="Bug reports">🐛</a> <a href="https://github.com/dictyBase/graphql-server/commits?author=cybersiddhu" title="Code">💻</a> <a href="#content-cybersiddhu" title="Content">🖋</a> <a href="https://github.com/dictyBase/graphql-server/commits?author=cybersiddhu" title="Documentation">📖</a> <a href="#maintenance-cybersiddhu" title="Maintenance">🚧</a></td>
    <td align="center"><a href="http://www.erichartline.net/"><img src="https://avatars3.githubusercontent.com/u/13489381?v=4" width="100px;" alt=""/><br /><sub><b>Eric Hartline</b></sub></a><br /><a href="https://github.com/dictyBase/graphql-server/issues?q=author%3Awildlifehexagon" title="Bug reports">🐛</a> <a href="https://github.com/dictyBase/graphql-server/commits?author=wildlifehexagon" title="Code">💻</a> <a href="#content-wildlifehexagon" title="Content">🖋</a> <a href="https://github.com/dictyBase/graphql-server/commits?author=wildlifehexagon" title="Documentation">📖</a> <a href="#maintenance-wildlifehexagon" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!