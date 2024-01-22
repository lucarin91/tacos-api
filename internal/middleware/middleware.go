package middleware

import (
	"log/slog"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/golang-jwt/jwt/v5"

	"github.com/lucarin91/tacos-api/internal/handler"
)

func Auth(secret []byte) func(next handler.UserIdHandler) httptreemux.HandlerFunc {
	return func(next handler.UserIdHandler) httptreemux.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			tokenStr := r.Header.Get("Authorization")

			token, err := jwt.ParseWithClaims(tokenStr, &handler.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})
			if err != nil {
				slog.Debug("Error parsing token", "error", err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(*handler.TokenClaims)
			if !token.Valid || !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next(w, r, params, claims.UserID)
		}
	}
}

func Log(next httptreemux.HandlerFunc) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		slog.Info("Request", "method", r.Method, "path", r.URL.Path)
		next(w, r, params)
	}
}
