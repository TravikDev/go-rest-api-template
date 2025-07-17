package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"go-rest-api-template/internal/auth"
)

type contextKey string

const UserIDKey contextKey = "userID"

// UserIDFromContext extracts the authenticated user ID from ctx.
func UserIDFromContext(ctx context.Context) (int, bool) {
	v := ctx.Value(UserIDKey)
	id, ok := v.(int)
	return id, ok
}

func JWTAuth(secret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseToken(tokenStr, secret)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		id, err := strconv.Atoi(claims.Sub)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserIDKey, id)
		next(w, r.WithContext(ctx))
	}
}
