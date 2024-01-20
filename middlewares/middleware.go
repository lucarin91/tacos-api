package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/golang-jwt/jwt/v5"

	"github.com/lucarin91/tacos-api/handlers"
)

func Auth(secret []byte) func(next handlers.UserIdHandler) httptreemux.HandlerFunc {
	return func(next handlers.UserIdHandler) httptreemux.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
			tokenStr := r.Header.Get("Authorization")

			token, err := jwt.ParseWithClaims(tokenStr, &handlers.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			})
			if err != nil {
				slog.Debug("Error parsing token", "error", err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(*handlers.TokenClaims)
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
