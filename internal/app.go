package internal

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"gophermart/internal/config"
	"gophermart/internal/domain/controllers/api/rest"
	"gophermart/internal/domain/controllers/api/rest/middleware"
	"gophermart/internal/domain/repositories"
	"os"
	"os/signal"
	"syscall"

	"context"
	"net/http"
	"time"
)

type App interface {
	Run() error
}

type AppImpl struct {
	s repositories.Storage
	h rest.RESTControllers
}

func New(s repositories.Storage, h rest.RESTControllers) *AppImpl {
	return &AppImpl{
		s: s,
		h: h,
	}
}

func (a AppImpl) Run() error {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	eg, errCtx := errgroup.WithContext(ctx)

	// Down migrations after stop app
	eg.Go(func() error {
		<-errCtx.Done()
		return a.s.Down(context.Background())
	})

	// Start server
	eg.Go(func() error {

		// Up migration
		if err := a.s.Up(errCtx); err != nil {
			return err
		}

		return a.listen(errCtx)
	})

	return eg.Wait()
}

func (a AppImpl) listen(ctx context.Context) error {

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	user := api.PathPrefix("/user").Subrouter()
	//auth := api.PathPrefix("").Subrouter()

	api.Use(middleware.RequestLogger)

	r.Handle("/ping", http.HandlerFunc(a.h.Health)).Methods(http.MethodGet)

	user.Handle("/register", http.HandlerFunc(a.h.UserRegister)).Methods(http.MethodPost)
	user.Handle("/login", http.HandlerFunc(a.h.UserLogin)).Methods(http.MethodPost)
	user.Handle("/orders", http.HandlerFunc(a.h.UserCreateOrders)).Methods(http.MethodPost)
	user.Handle("/orders", http.HandlerFunc(a.h.UserGetOrders)).Methods(http.MethodGet)
	user.Handle("/balance", http.HandlerFunc(a.h.UserBalance)).Methods(http.MethodGet)
	user.Handle("/balance/withdraw", http.HandlerFunc(a.h.UserBalanceWithdraw)).Methods(http.MethodPost)
	user.Handle("/withdrawals", http.HandlerFunc(a.h.UserWithdraws)).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Conf.ServerAddr(),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	go func() {
		<-ctx.Done()
		fmt.Println("Stop listen")
		_ = srv.Shutdown(context.Background())
	}()

	fmt.Println("Start listen")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
