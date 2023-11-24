package rest

import (
	"gophermart/internal/domain/controllers/api/rest/dto"
	"net/http"
)

func (h RESTControllersImpl) UserRegister(w http.ResponseWriter, r *http.Request) {
	_ = h.userService.Register(dto.RegisterUser{
		Login:    "",
		Password: "",
	})

	w.WriteHeader(http.StatusOK)
}
