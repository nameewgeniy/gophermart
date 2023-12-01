package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func (h RESTControllersImpl) UserWithdraws(w http.ResponseWriter, r *http.Request) {

	res, err := h.userService.UserWithdraws(uuid.UUID{})

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
