package handlers

import "net/http"

func (h MuxHandlers) UserWithdrawsHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
