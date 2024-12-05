// internal/models/client_test.go
package models

import (
	"testing"
)

func TestClientCreation(t *testing.T) {
	client := Client{
		ID:         1,
		LastName:   "Иванов",
		FirstName:  "Иван",
		MiddleName: "Иванович",
		Phone:      "81234567890",
	}

	if client.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", client.ID)
	}

	if client.LastName != "Иванов" {
		t.Errorf("expected LastName to be 'Иванов', got %s", client.LastName)
	}

	if client.FirstName != "Иван" {
		t.Errorf("expected FirstName to be 'Иван', got %s", client.FirstName)
	}

	if client.MiddleName != "Иванович" {
		t.Errorf("expected MiddleName to be 'Иванович', got %s", client.MiddleName)
	}

	if client.Phone != "81234567890" {
		t.Errorf("expected Phone to be '81234567890', got %s", client.Phone)
	}
}
