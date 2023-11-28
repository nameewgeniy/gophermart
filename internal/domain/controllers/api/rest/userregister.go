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

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Authorization", "dfsdfsdf")
	w.WriteHeader(http.StatusOK)
}
