package handlers

import "net/http"

func (h MuxHandlers) UserGetOrdersHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}