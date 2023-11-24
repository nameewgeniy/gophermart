package rest

import "net/http"

func (h RESTControllersImpl) UserGetOrders(w http.ResponseWriter, r *http.Request) {

	h.userService.UserGetOrders()
	w.WriteHeader(http.StatusOK)
}
