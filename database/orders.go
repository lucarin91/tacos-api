package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/lucarin91/tacos-api/models"
)

func GetOrders(ctx context.Context, pool *pgxpool.Pool, userID uuid.UUID) ([]models.Order, error) {
	rows, err := pool.Query(ctx, `
		select orders.id, ingredients.id, ingredients.name, ingredients.category
		from orders
		inner join orders_ingredients on orders.id = orders_ingredients.order_id
		inner join ingredients on orders_ingredients.ingredient_id = ingredients.id 
		where user_id = $1`,
		userID)
	if err != nil {
		return nil, err
	}

	ordersMap := make(map[uuid.UUID]models.Order)
	for rows.Next() {
		var orderID uuid.UUID
		var ingredient models.Ingredient
		err := rows.Scan(&orderID, &ingredient.ID, &ingredient.Name, &ingredient.Category)
		if err != nil {
			return nil, err
		}

		order, ok := ordersMap[orderID]
		if !ok {
			order = models.Order{
				ID: orderID,
			}
		}
		order.Ingredients = append(order.Ingredients, ingredient)
		ordersMap[orderID] = order
	}

	orders := make([]models.Order, 0, len(ordersMap))
	for _, order := range ordersMap {
		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrder(ctx context.Context, pool *pgxpool.Pool, userID, id uuid.UUID) (*models.Order, error) {
	rows, err := pool.Query(ctx, `
		select ingredients.id, ingredients.name, ingredients.category
		from orders
		inner join orders_ingredients on orders.id = orders_ingredients.order_id
		inner join ingredients on orders_ingredients.ingredient_id = ingredients.id 
		where orders.user_id = $1 and orders.id = $2`,
		userID, id)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	order := &models.Order{ID: id}
	for rows.Next() {
		var ingredient models.Ingredient
		err := rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Category)
		if err != nil {
			return nil, err
		}

		order.Ingredients = append(order.Ingredients, ingredient)
	}

	return order, nil
}

func CreateOrder(ctx context.Context, pool *pgxpool.Pool, userID uuid.UUID, order models.Order) error {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	for _, ingredient := range order.Ingredients {
		_, err := tx.Exec(
			ctx,
			"insert into orders_ingredients (order_id, ingredient_id) values ($1, $2)",
			order.ID, ingredient.ID,
		)
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx, "insert into orders (id, user_id) values ($1, $2)", order.ID, userID)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func DeleteOrder(ctx context.Context, pool *pgxpool.Pool, userID, id uuid.UUID) error {
	_, err := pool.Exec(ctx, "delete from orders where user_id = $1 and id = $2", userID, id)
	if err != nil {
		return err
	}

	return nil
}
