package rest

import (
	"gophermart/internal/domain/controllers/api/rest/dto"
	"net/http"
)

func (h RESTControllersImpl) UserCreateOrders(w http.ResponseWriter, r *http.Request) {

	h.userService.UserCreateOrders(dto.CreateOrder{})
	w.WriteHeader(http.StatusOK)
}
