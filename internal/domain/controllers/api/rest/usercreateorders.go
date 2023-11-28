package rest

import (
	"gophermart/internal/domain/controllers/api/rest/dto"
	"net/http"
)

func (h RESTControllersImpl) UserCreateOrders(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/plain; charset=utf-8")

	h.userService.UserCreateOrders(dto.CreateOrder{})

	w.Write([]byte("weqrwrwe"))

	w.WriteHeader(http.StatusAccepted)
}
