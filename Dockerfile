FROM golang:latest AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/server.go

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/server /app/server

USER 65532:65532 # nobody 
EXPOSE 8080

ENTRYPOINT ["/app/server"]