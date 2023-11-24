package rest

import "net/http"

func (h RESTControllersImpl) UserBalance(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
