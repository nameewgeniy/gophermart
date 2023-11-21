package main

import (
	"gophermart/internal/config"
	"gophermart/internal/logger"
	"gophermart/internal/server"
	"gophermart/internal/server/handlers"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	// Init configuration
	config.Singleton()

	// Init logger
	if err := logger.Singleton(config.Conf.LogLevel()); err != nil {
		return err
	}

	muxHandlers := handlers.NewMuxHandlers()

	srv := server.New(muxHandlers)

	if err := srv.Run(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
