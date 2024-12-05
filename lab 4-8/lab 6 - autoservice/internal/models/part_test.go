// internal/models/part_test.go
package models

import (
	"testing"
)

func TestPartCreation(t *testing.T) {
	part := Part{
		ID:          1,
		Name:        "Шина",
		Description: "Летняя шина",
		Price:       5000.00,
	}

	if part.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", part.ID)
	}

	if part.Name != "Шина" {
		t.Errorf("expected Name to be 'Шина', got %s", part.Name)
	}

	if part.Description != "Летняя шина" {
		t.Errorf("expected Description to be 'Летняя шина', got %s", part.Description)
	}

	if part.Price != 5000.00 {
		t.Errorf("expected Price to be 5000.00, got %f", part.Price)
	}
}
