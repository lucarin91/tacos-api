package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/lucarin91/tacos-api/database"
	"github.com/lucarin91/tacos-api/models"
)

func GetIngredients(pool *pgxpool.Pool) httptreemux.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string) {
		ingredients, err := database.GetIngredients(r.Context(), pool)
		if err != nil {
			models.WriteInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(ingredients)
		if err != nil {
			models.WriteInternalError(w, err)
			return
		}
	}
}
