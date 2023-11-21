package handlers

import "net/http"

func (h MuxHandlers) UserRegisterHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
