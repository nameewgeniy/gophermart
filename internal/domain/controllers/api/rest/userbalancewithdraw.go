package rest

import (
	"encoding/json"
	"gophermart/internal/domain/controllers/api/rest/dto"
	"io"
	"net/http"
)

// 200 — успешная обработка запроса;
// 401 — пользователь не авторизован;
// 402 — на счету недостаточно средств;
// 422 — неверный номер заказа;
// 500 — внутренняя ошибка сервера.

func (h RESTControllersImpl) UserBalanceWithdraw(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var uDto dto.UserBalanceWithdraw
	if err = json.Unmarshal(bytes, &uDto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = h.userService.UserBalanceWithdraw(uDto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
