package handlers

import "net/http"

func (h MuxHandlers) UserLoginHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
