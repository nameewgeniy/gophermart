package handlers

import (
	"net/http"
)

func (h MuxHandlers) PingHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
