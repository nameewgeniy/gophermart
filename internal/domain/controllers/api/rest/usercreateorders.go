package rest

import (
	"gophermart/internal/domain/controllers/api/rest/dto"
	"io"
	"net/http"
)

//200 — номер заказа уже был загружен этим пользователем;
//202 — новый номер заказа принят в обработку;
//400 — неверный формат запроса;
//401 — пользователь не аутентифицирован;
//409 — номер заказа уже был загружен другим пользователем;
//422 — неверный формат номера заказа;
//500 — внутренняя ошибка сервера.

func (h RESTControllersImpl) UserCreateOrders(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.UserCreateOrders(dto.CreateOrder{
		Number: string(bytes),
	})

	if err != nil {
		// TODO обработать номер заказа уже был загружен этим пользователем;
		// 	номер заказа уже был загружен другим пользователем
		// 	неверный формат номера заказа
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusAccepted)
}
