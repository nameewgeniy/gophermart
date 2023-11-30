package rest

import (
	"encoding/json"
	"gophermart/internal/domain/controllers/api/rest/dto"
	"io"
	"net/http"
)

func (h RESTControllersImpl) UserLogin(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var uDto dto.LoginUser
	if err = json.Unmarshal(bytes, &uDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := h.userService.Login(uDto)

	if err != nil {
		// TODO неверная пара логин/пароль; 401
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", res.AccessToken)
	w.WriteHeader(http.StatusOK)
}
