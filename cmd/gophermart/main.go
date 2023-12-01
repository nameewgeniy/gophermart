package main

import (
	"gophermart/internal"
	"gophermart/internal/config"
	"gophermart/internal/domain/controllers/api/rest"
	"gophermart/internal/domain/repositories/pg"
	"gophermart/internal/domain/services/auth"
	"gophermart/internal/domain/services/user"
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
	config.New().Parse()

	// Init logger
	if err := logger.Singleton(config.Instance.LogLevel()); err != nil {
		return err
	}

	// Init pg storage
	storage, err := pg.New()
	if err != nil {
		return err
	}

	// Init services
	au := auth.NewAuthService(
		storage,
		config.Instance.AuthSecretKey(),
		config.Instance.AuthAccessTTL(),
		config.Instance.AuthRefreshTTL(),
	)

	us := user.NewUserService(storage, storage, storage, au)

	// Init controllers
	restApi := rest.New(us, au)

	// Init app
	app := internal.New(
		storage,
		restApi,
	)

	return app.Run()
}
