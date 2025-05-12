# Level 3


## Goals

- Add a simple logger to the server and client.
- Add a simple error handler to the server and client.

### Logging 

- See [zerolog](https://github.com/rs/zerolog)

```bash

go get github.com/rs/zerolog/log

```

### Middlewares

- See [middlewares](https://docs.gofiber.io/category/-middleware/) and [contrib](https://docs.gofiber.io/contrib/)
    - Logger
    - ContextLogger. Have a look at zerolog's `logger.WithContext(ctx)`
    - ErrorHandler