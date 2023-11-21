package handlers

import "net/http"

func (h MuxHandlers) UserCreateOrdersHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
