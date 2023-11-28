package rest

import "net/http"

func (h RESTControllersImpl) UserWithdraws(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
