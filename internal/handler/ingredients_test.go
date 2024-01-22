package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/lucarin91/tacos-api/internal/handler"
	"github.com/lucarin91/tacos-api/internal/model"
)

func GetDBPool(t *testing.T) *pgxpool.Pool {
	pool, err := pgxpool.New(context.TODO(), "postgresql://postgres:postgres@localhost:5432")
	if err != nil {
		t.Fatal(err)
	}

	return pool
}

func TestGetIngredients(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()

	want := []model.Ingredient{
		{
			ID:       uuid1,
			Name:     "value1",
			Category: "category1",
		},
		{
			ID:       uuid2,
			Name:     "value2",
			Category: "category2",
		},
	}
	wantJson, err := json.Marshal(want)
	if err != nil {
		t.Fatal(err)
	}

	pool := GetDBPool(t)

	_, err = pool.Exec(context.TODO(), `truncate table ingredients`)
	if err != nil {
		t.Fatal(err)
	}

	_, err = pool.Exec(context.TODO(), `
		insert into ingredients (id, name, category)
		values ($1, 'value1', 'category1'),
		       ($2, 'value2', 'category2')
	`, uuid1, uuid2)
	if err != nil {
		t.Fatal(err)
	}

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/ingredients", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call handler
	rr := httptest.NewRecorder()
	handler := handler.GetIngredients(pool)
	handler(rr, req, nil)

	// checks
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code: got %v want %v", status, http.StatusOK)
	}
	got := strings.TrimSpace(rr.Body.String())
	if got != string(wantJson) {
		t.Errorf("body: got %v, want %v", rr.Body.String(), string(wantJson))
	}
}
