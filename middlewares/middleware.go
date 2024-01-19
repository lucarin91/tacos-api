package middlewares

import (
	"log/slog"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"

	"github.com/lucarin91/tacos-api/handlers"
)

func Auth(next handlers.UserIdHandler) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		// TODO: implement auth.
		// authorize user and inject user id.
		userID := uuid.MustParse("6ba7b814-9dad-11d1-80b4-00c04fd430c9")
		next(w, r, params, userID)
	}
}

func Log(next httptreemux.HandlerFunc) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		slog.Info("Request", "method", r.Method, "path", r.URL.Path)
		next(w, r, params)
	}
}
