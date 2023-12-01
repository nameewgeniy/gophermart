package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (h RESTControllersImpl) UserGetOrders(w http.ResponseWriter, r *http.Request) {

	res, err := h.userService.UserGetOrders(uuid.UUID{}) // Идентиыикатор авторизованного прользователя

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
