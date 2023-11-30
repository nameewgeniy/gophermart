package rest

import (
	"encoding/json"
	"net/http"
)

func (h RESTControllersImpl) UserWithdraws(w http.ResponseWriter, r *http.Request) {

	res, err := h.userService.UserWithdraws()

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
