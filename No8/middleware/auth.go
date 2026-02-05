package middleware

import (
	"context"
	"net/http"
	"strings"
)

type ctxKey string

const UserIDKey ctxKey = "user_id"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		userID := parts[1]
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
