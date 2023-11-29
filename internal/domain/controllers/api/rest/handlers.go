package rest

import (
	"github.com/gorilla/mux"
	"gophermart/internal/domain/controllers/api/rest/middleware"
	"gophermart/internal/domain/services"
	"gophermart/internal/domain/services/auth"
	"net/http"
)

type RESTControllers interface {
	Health(http.ResponseWriter, *http.Request)

	UserRegister(http.ResponseWriter, *http.Request)
	UserLogin(http.ResponseWriter, *http.Request)

	UserCreateOrders(http.ResponseWriter, *http.Request)
	UserGetOrders(http.ResponseWriter, *http.Request)

	UserBalance(http.ResponseWriter, *http.Request)
	UserBalanceWithdraw(http.ResponseWriter, *http.Request)
	UserWithdraws(http.ResponseWriter, *http.Request)

	Router() *mux.Router
}

type RESTControllersImpl struct {
	userService services.UserService
	authService auth.AuthService
}

func New(us services.UserService, as auth.AuthService) *RESTControllersImpl {
	return &RESTControllersImpl{
		userService: us,
		authService: as,
	}
}

func (h RESTControllersImpl) Router() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	user := api.PathPrefix("/user").Subrouter()
	rAuth := api.PathPrefix("").Subrouter()

	api.Use(middleware.RequestLogger)
	rAuth.Use(middleware.Auth)

	rAuth.HandleFunc("/health", h.Health).Methods(http.MethodGet)

	user.HandleFunc("/register", h.UserRegister).Methods(http.MethodPost)
	user.HandleFunc("/login", h.UserLogin).Methods(http.MethodPost)
	user.HandleFunc("/orders", h.UserCreateOrders).Methods(http.MethodPost)
	user.HandleFunc("/orders", h.UserGetOrders).Methods(http.MethodGet)
	user.HandleFunc("/balance", h.UserBalance).Methods(http.MethodGet)
	user.HandleFunc("/balance/withdraw", h.UserBalanceWithdraw).Methods(http.MethodPost)
	user.HandleFunc("/withdrawals", h.UserWithdraws).Methods(http.MethodGet)

	return r
}
