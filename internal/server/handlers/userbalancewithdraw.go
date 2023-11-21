package handlers

import "net/http"

func (h MuxHandlers) UserBalanceWithdrawHandle(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}
