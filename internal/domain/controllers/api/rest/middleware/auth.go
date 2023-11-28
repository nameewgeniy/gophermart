package middleware

import (
	"net/http"
)

// Auth — middleware авторизации
func Auth(next http.HandlerFunc) http.HandlerFunc {

	logFn := func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)

	}
	return logFn
}
