package main

import (
	"fmt"
	"gophermart/internal/config"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	// Init configuration
	config.Singleton()

	srv := &http.Server{
		Addr:         config.Conf.ServerAddr(),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	//go func() {
	//	<-ctx.Done()
	//	_ = srv.Shutdown(context.Background())
	//}()

	fmt.Println("Server up")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
