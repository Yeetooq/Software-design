// internal/handlers/order_handler_test.go
package handlers

import (
	"autoservice/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	order := models.Order{ClientID: 1, PartID: 1, Quantity: 2, TotalCost: 200.00}

	mock.ExpectExec("INSERT INTO orders").WithArgs(order.ClientID, order.PartID, order.Quantity, order.TotalCost).WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := addOrder(db, order)
	if err != nil {
		t.Errorf("error was not expected while adding order: %s", err)
	}

	if id != 1 {
		t.Errorf("expected id to be 1, but got %d", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM orders").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = deleteOrder(db, 1)
	if err != nil {
		t.Errorf("error was not expected while deleting order: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
