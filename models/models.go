package models

import (
	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID    `json:"id"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
}
