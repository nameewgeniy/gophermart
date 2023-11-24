package rest

import (
	"gophermart/internal/domain/services"
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
}

type RESTControllersImpl struct {
	userService services.UserService
	authService services.AuthService
}

func New(us services.UserService, as services.AuthService) *RESTControllersImpl {
	return &RESTControllersImpl{
		userService: us,
		authService: as,
	}
}
