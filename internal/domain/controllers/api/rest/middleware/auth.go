package middleware

import (
	"gophermart/internal/domain/models"
	"gophermart/internal/domain/services/auth"
	"net/http"
)

// Auth — middleware авторизации
func Auth(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		if auth.Instance == nil {
			panic("Auth not instance")
		}

		if _, err := auth.Instance.VerifyToken(r.Header.Get("Authorization"), models.TokenTypeAccess); err != nil {
			http.Error(w, auth.ErrInvalidToken.Error(), http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	}
	return http.HandlerFunc(fn)
}
