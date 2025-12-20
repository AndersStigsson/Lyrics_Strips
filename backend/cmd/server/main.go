package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var interruptSignals = []os.Signal{
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		interruptSignals...,
	)
	defer stop()

	config, err := NewConfigFromEnv()
	if err != nil {
		slog.Error("Unable to load config", "error", err)

		return
	}

	loggerOpts := &slog.HandlerOptions{}
	if config.debug {
		loggerOpts = &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}
	}

	logger := slog.New(
		slog.NewTextHandler(os.Stdout, loggerOpts),
	)

	server := NewServer(config, logger)
	server.Start(ctx)
}
