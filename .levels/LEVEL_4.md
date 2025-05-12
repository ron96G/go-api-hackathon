# Level 4

## Goals

- Add more middlewares including:
    - Authentication using OAuth2.0 "client credentials" token flow
    - Input validation based on the openapi spec
- Implement something more complex than just a simple CRUD operation

### Auth Middleware

- Implement a JWT authentication middleware, see [jwt-mw](https://docs.gofiber.io/contrib/jwt/)

```bash
export TRUSTED_ISSUER="https://iris-playground.live.dhei.telekom.de/auth/realms/default"
# only used for client
export TOKEN_URL="https://iris-playground.live.dhei.telekom.de/auth/realms/default/protocol/openid-connect/token"
export CLIENT_ID="go-api-hackathon-client"
export CLIENT_SECRET="topsecret"

```

- Add the login to your client and send the JWT token in the header of your requests
    - See [clientcredentials](https://pkg.go.dev/golang.org/x/oauth2/clientcredentials)
    - See [dotenv](https://github.com/joho/godotenv) to not hardcode the credentials in the code.

### Input Validation

- See [validator-mw](github.com/oapi-codegen/fiber-middleware)

### More Complex Logic
