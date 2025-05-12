package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/ron96G/go-api-hackathon/internal/api"
	"github.com/ron96G/go-api-hackathon/internal/service"
)

var (
	rootCtx = context.Background()
)

func main() {
	ctx, cancel := signal.NotifyContext(rootCtx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := api.NewStrictHandler(service.NewService(ctx), nil)
	app := fiber.New(fiber.Config{
		ReadTimeout: 2 * time.Second,
	})
	api.RegisterHandlers(app, server)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			panic(err)
		}
	}()

	<-ctx.Done()
	if err := app.Shutdown(); err != nil {
		panic(err)
	}
}
