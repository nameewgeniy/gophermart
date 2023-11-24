package rest

import "net/http"

func (h RESTControllersImpl) UserLogin(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
