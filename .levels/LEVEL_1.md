# Level 1

## Goals

- Define OpenAPI File
- Setup basic Go Project
- Generate Server Code from OpenAPI
- Implement & Run Basic Server
- Test with curl/Browser

### Define OpenAPI File

- Define a "GET", "GET {id}", "POST" and "DELETE {id}" endpoint
- See [OpenAPI Specification](https://spec.openapis.org/oas/v3.1.0#)
- See [Editor UI](https://ron96g.github.io/oas-linter/?ruleset=oas&schema=openapi.v3.0)
- **Fallback:** See [Petstore Example](../api/api.yaml)

### Setup basic Go Project

- See [Go Project Layout](https://github.com/golang-standards/project-layout)

```bash
# navigate to your project directory

go mod init github.com/<your_username>/<your_project_name>
mkdir -p cmd internal/api pkg tools
touch cmd/main.go tools/tools.go
# Add some code to main.go
go run cmd/main.go
```

### Generate Server Code from OpenAPI

- See [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)
- See [how-to-configure](https://github.com/oapi-codegen/oapi-codegen?tab=readme-ov-file#generating-server-side-boilerplate)
- Configure it to generate a fiber server, see [fiber](https://github.com/gofiber/fiber/tree/v2.52.6)

```bash
go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1

# Add the config to generate the server code, see provided links

# generate the server code
go generate tools/tools.go

# get the needed dependencies
go mod tidy

```

### Implement & Run Basic Server

- See [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen)

```bash
mkdir -p cmd/server 
touch cmd/server/server.go
```

### Test with curl/Browser 

```bash
go run cmd/server/server.go

# Test with curl
curl http://localhost:8080/ping

```