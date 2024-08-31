package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/guluzadehh/kode_test/utils"
	"golang.org/x/crypto/bcrypt"
)

type contextKey string

const userContextKey contextKey = "user"

func (h *Handler) BasicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			utils.ErrorJSON(w, http.StatusUnauthorized, fmt.Errorf("invalid authentication format"))
			return
		}

		user, exists := h.sr.Store.GetByUsername(username)
		if !exists || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			utils.ErrorJSON(w, http.StatusUnauthorized, fmt.Errorf("wrong credentials"))
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
