package main

import (
	"gophermart/internal"
	"gophermart/internal/config"
	"gophermart/internal/domain/controllers/api/rest"
	"gophermart/internal/domain/repositories/pg"
	"gophermart/internal/domain/services"
	"gophermart/internal/logger"
	"log"
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

	// Init pg storage
	storage, err := pg.New()
	if err != nil {
		return err
	}

	// Init services
	au := services.NewAuthService(
		storage,
		config.Conf.AuthSecretKey(),
		config.Conf.AuthAccessTTL(),
		config.Conf.AuthRefreshTTL(),
	)

	us := services.NewUserService(storage)

	// Init controllers
	restApi := rest.New(us, au)

	// Init app
	app := internal.New(
		storage,
		restApi,
	)

	return app.Run()
}
