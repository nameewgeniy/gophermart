package handlers

import "net/http"

type Handlers interface {
	PingHandle(http.ResponseWriter, *http.Request)

	UserRegisterHandle(http.ResponseWriter, *http.Request)
	UserLoginHandle(http.ResponseWriter, *http.Request)

	UserCreateOrdersHandle(http.ResponseWriter, *http.Request)
	UserGetOrdersHandle(http.ResponseWriter, *http.Request)

	UserBalanceHandle(http.ResponseWriter, *http.Request)
	UserBalanceWithdrawHandle(http.ResponseWriter, *http.Request)
	UserWithdrawsHandle(http.ResponseWriter, *http.Request)
}

type MuxHandlers struct {
}

func NewMuxHandlers() *MuxHandlers {
	return &MuxHandlers{}
}
