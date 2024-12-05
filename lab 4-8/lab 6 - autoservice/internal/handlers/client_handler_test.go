// internal/handlers/client_handler_test.go
package handlers

import (
	"autoservice/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	client := models.Client{LastName: "Doe", FirstName: "John", MiddleName: "Smith", Phone: "81234567890"}

	mock.ExpectExec("INSERT INTO clients").WithArgs(client.LastName, client.FirstName, client.MiddleName, client.Phone).WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := addClient(db, client)
	if err != nil {
		t.Errorf("error was not expected while adding client: %s", err)
	}

	if id != 1 {
		t.Errorf("expected id to be 1, but got %d", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM clients").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = deleteClient(db, 1)
	if err != nil {
		t.Errorf("error was not expected while deleting client: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
