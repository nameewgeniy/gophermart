package rest

import (
	"encoding/json"
	"gophermart/internal/domain/controllers/api/rest/dto"
	"io"
	"net/http"
)

func (h RESTControllersImpl) UserRegister(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var uDto dto.RegisterUser
	if err = json.Unmarshal(bytes, &uDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.userService.Register(uDto)

	if err != nil {
		// TODO обработать 409 - логин уже занят
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.userService.Login(dto.LoginUser{
		Login:    uDto.Login,
		Password: uDto.Password,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", res.AccessToken)
	w.WriteHeader(http.StatusOK)
}
