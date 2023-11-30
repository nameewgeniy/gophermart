package rest

import (
	"encoding/json"
	"net/http"
)

func (h RESTControllersImpl) UserGetOrders(w http.ResponseWriter, r *http.Request) {

	res, err := h.userService.UserGetOrders()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")

	w.Write(body)
	w.WriteHeader(http.StatusOK)
}
