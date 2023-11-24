package rest

import (
	"net/http"
)

func (h RESTControllersImpl) Health(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
