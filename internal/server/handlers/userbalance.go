package handlers

import "net/http"

func (h MuxHandlers) UserBalanceHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
