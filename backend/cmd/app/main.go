package main

import (
	"backend/cmd/server"
	"context"
	"log/slog"
	"os"
	"os/signal"
)

var logger slog.Logger

func main() {
	app := server.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Run(ctx)
	if err != nil {
		logger.Error("Failed to start..")
	}
}
