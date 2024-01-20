package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/lucarin91/tacos-api/models"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id"`
}

func GetToken(secret []byte) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		params := r.URL.Query()
		username := params.Get("username")
		password := params.Get("password")

		// FIXME: implement real authentification.
		if username != "admin" || password != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			&TokenClaims{
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
				},
				// FIXME: return a real user ID.
				UserID: uuid.MustParse("6ba7b814-9dad-11d1-80b4-00c04fd430c9"),
			})

		tokenStr, err := token.SignedString(secret)
		if err != nil {
			models.WriteInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(models.Token{Token: tokenStr})
		if err != nil {
			models.WriteInternalError(w, err)
			return
		}
	}
}
