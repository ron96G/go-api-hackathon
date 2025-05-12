# Level 2

## Goals

- Implemnt a simple Go client that interacts with the OpenAPI server.
- Use the OpenAPI spec to generate a Go client.
- Test the client-server interaction.

### Generate Go Client from 

- See Level 1 for the OpenAPI spec and how to generate the Go client.

### Implement Minimal CLI Client

- Use the generated Go client and use it to implement a minimal CLI client.
- You may have a look at [cobra](https://cobra.dev/#about) but it's probably too much for this level (Go has a lot of good libs for CLIs).

### Test Client-Server Interaction 

```bash

go run cmd/client/client.go --get --id something # or however you want to implement it

# you can also install the executable and run it from anywhere
GOBIN=~/.local/bin go install cmd/client/client.go
client --get --id something
```