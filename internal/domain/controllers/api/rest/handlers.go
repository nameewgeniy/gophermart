package rest

import (
	"github.com/gorilla/mux"
	"gophermart/internal/domain/controllers/api/rest/middleware"
	"gophermart/internal/domain/services/auth"
	"gophermart/internal/domain/services/user"
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
	userService user.UserService
	authService auth.AuthService
}

func New(us user.UserService, as auth.AuthService) *RESTControllersImpl {
	return &RESTControllersImpl{
		userService: us,
		authService: as,
	}
}

func (h RESTControllersImpl) Router() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	usr := api.PathPrefix("/user").Subrouter()
	usrAuth := api.PathPrefix("/user").Subrouter()

	api.Use(middleware.RequestLogger)
	usrAuth.Use(middleware.Auth)

	r.HandleFunc("/health", h.Health).Methods(http.MethodGet)

	usr.HandleFunc("/register", h.UserRegister).Methods(http.MethodPost)
	usr.HandleFunc("/login", h.UserLogin).Methods(http.MethodPost)

	usrAuth.HandleFunc("/orders", h.UserCreateOrders).Methods(http.MethodPost)
	usrAuth.HandleFunc("/orders", h.UserGetOrders).Methods(http.MethodGet)
	usrAuth.HandleFunc("/balance", h.UserBalance).Methods(http.MethodGet)
	usrAuth.HandleFunc("/balance/withdraw", h.UserBalanceWithdraw).Methods(http.MethodPost)
	usrAuth.HandleFunc("/withdrawals", h.UserWithdraws).Methods(http.MethodGet)

	return r
}
