package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/lucarin91/tacos-api/internal/database"
	"github.com/lucarin91/tacos-api/internal/model"
)

type UserIdHandler func(w http.ResponseWriter, r *http.Request, params map[string]string, userID uuid.UUID)

func GetOrders(pool *pgxpool.Pool) UserIdHandler {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string, userID uuid.UUID) {
		orders, err := database.GetOrders(r.Context(), pool, userID)
		if err != nil {
			model.WriteInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(orders)
		if err != nil {
			model.WriteInternalError(w, err)
			return
		}
	}
}

func GetOrder(pool *pgxpool.Pool) UserIdHandler {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string, userID uuid.UUID) {
		id, err := uuid.Parse(params["id"])
		if err != nil {
			model.WriteClientError(w, err)
			return
		}

		order, err := database.GetOrder(r.Context(), pool, userID, id)
		if err != nil {
			model.WriteClientError(w, err)
		}

		if order == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(order)
		if err != nil {
			model.WriteInternalError(w, err)
			return
		}
	}
}

func validateIngredients(ingredients []model.Ingredient) error {
	protein, bread := false, false

	for _, ingredient := range ingredients {
		if ingredient.ID == uuid.Nil {
			return fmt.Errorf("ingredient id cannot be nil")
		}
		if len(ingredient.Name) == 0 {
			return fmt.Errorf("ingredient name cannot be nil")
		}

		if ingredient.Category == "protein" {
			protein = true
		}
		if ingredient.Category == "bread" {
			bread = true
		}
	}

	if !protein || !bread {
		return fmt.Errorf("order must have a bread and a protein")
	}

	return nil
}

func CreateOrder(pool *pgxpool.Pool) UserIdHandler {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string, userID uuid.UUID) {
		var in struct {
			Ingredients []model.Ingredient `json:"ingredients"`
		}
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			model.WriteClientError(w, err)
			return
		}

		err = validateIngredients(in.Ingredients)
		if err != nil {
			model.WriteClientError(w, err)
			return
		}

		order := model.Order{
			ID:          uuid.New(),
			Ingredients: in.Ingredients,
		}

		err = database.CreateOrder(r.Context(), pool, userID, order)
		if err != nil {
			model.WriteInternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteOrder(pool *pgxpool.Pool) UserIdHandler {
	return func(w http.ResponseWriter, r *http.Request, params map[string]string, userID uuid.UUID) {
		id, err := uuid.Parse(params["id"])
		if err != nil {
			model.WriteClientError(w, err)
			return
		}

		err = database.DeleteOrder(r.Context(), pool, userID, id)
		if err != nil {
			model.WriteInternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
