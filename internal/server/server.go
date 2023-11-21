package server

import (
	"github.com/gorilla/mux"
	"gophermart/internal/config"
	"gophermart/internal/server/handlers"
	"gophermart/internal/server/handlers/middleware"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Run() error
}

type Srv struct {
	h handlers.Handlers
}

func New(h handlers.Handlers) *Srv {
	return &Srv{
		h: h,
	}
}

func (s Srv) Run() error {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	errorCh := make(chan error)
	defer close(errorCh)

	return s.listen(ctx)
}

func (s Srv) listen(ctx context.Context) error {

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	user := api.PathPrefix("/user").Subrouter()

	r.Handle("/ping", middleware.RequestLogger(s.h.PingHandle)).Methods(http.MethodGet)

	user.Handle("/register", middleware.RequestLogger(s.h.UserRegisterHandle)).Methods(http.MethodPost)
	user.Handle("/login", middleware.RequestLogger(s.h.UserLoginHandle)).Methods(http.MethodPost)
	user.Handle("/orders", middleware.RequestLogger(s.h.UserCreateOrdersHandle)).Methods(http.MethodPost)
	user.Handle("/orders", middleware.RequestLogger(s.h.UserGetOrdersHandle)).Methods(http.MethodGet)
	user.Handle("/balance", middleware.RequestLogger(s.h.UserBalanceHandle)).Methods(http.MethodGet)
	user.Handle("/balance/withdraw", middleware.RequestLogger(s.h.UserBalanceWithdrawHandle)).Methods(http.MethodPost)
	user.Handle("/withdrawals", middleware.RequestLogger(s.h.UserWithdrawsHandle)).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         config.Conf.ServerAddr(),
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	go func() {
		<-ctx.Done()
		_ = srv.Shutdown(context.Background())
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
