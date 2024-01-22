package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/lucarin91/tacos-api/internal/model"
)

func GetIngredients(ctx context.Context, pool *pgxpool.Pool) ([]model.Ingredient, error) {
	rows, err := pool.Query(ctx, `select id, name, category from ingredients`)
	if err != nil {
		return nil, err
	}

	var ingredients []model.Ingredient
	for rows.Next() {
		var ingredient model.Ingredient
		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Category)
		if err != nil {
			return nil, err
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients, nil
}
