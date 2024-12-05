// internal/handlers/part_handler_test.go
package handlers

import (
	"autoservice/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddPart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	part := models.Part{Name: "Wheel", Description: "Car wheel", Price: 100.00}

	mock.ExpectExec("INSERT INTO parts").WithArgs(part.Name, part.Description, part.Price).WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := addPart(db, part)
	if err != nil {
		t.Errorf("error was not expected while adding part: %s", err)
	}

	if id != 1 {
		t.Errorf("expected id to be 1, but got %d", id)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletePart(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM parts").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = deletePart(db, 1)
	if err != nil {
		t.Errorf("error was not expected while deleting part: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
