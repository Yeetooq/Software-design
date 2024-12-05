// internal/models/order_test.go
package models

import (
	"testing"
)

func TestOrderCreation(t *testing.T) {
	order := Order{
		ID:        1,
		ClientID:  1,
		PartID:    1,
		Quantity:  2,
		TotalCost: 10000.00,
	}

	if order.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", order.ID)
	}

	if order.ClientID != 1 {
		t.Errorf("expected ClientID to be 1, got %d", order.ClientID)
	}

	if order.PartID != 1 {
		t.Errorf("expected PartID to be 1, got %d", order.PartID)
	}

	if order.Quantity != 2 {
		t.Errorf("expected Quantity to be 2, got %d", order.Quantity)
	}

	if order.TotalCost != 10000.00 {
		t.Errorf("expected TotalCost to be 10000.00, got %f", order.TotalCost)
	}
}
