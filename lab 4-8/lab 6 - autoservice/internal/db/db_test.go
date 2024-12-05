// internal/db/db_test.go
package db

import (
	"autoservice/internal/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestOpen(t *testing.T) {
	db, err := Open()
	if err != nil {
		t.Errorf("error was not expected while opening database: %s", err)
	}
	defer db.Close()

	if db == nil {
		t.Error("expected database connection, but got nil")
	}
}

func TestCreateTables(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS clients").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS parts").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS orders").WillReturnResult(sqlmock.NewResult(0, 0))

	CreateTables(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestClearDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("DELETE FROM clients").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("DELETE FROM parts").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("DELETE FROM orders").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("DELETE FROM sqlite_sequence").WithArgs("clients").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("DELETE FROM sqlite_sequence").WithArgs("parts").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("DELETE FROM sqlite_sequence").WithArgs("orders").WillReturnResult(sqlmock.NewResult(0, 0))

	ClearDatabase(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPopulateDatabase(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	clients := []models.Client{
		{LastName: "Иванов", FirstName: "Иван", MiddleName: "Иванович", Phone: "81234567890"},
		{LastName: "Петров", FirstName: "Петр", MiddleName: "Петрович", Phone: "80987654321"},
		{LastName: "Сидоров", FirstName: "Сидор", MiddleName: "Сидорович", Phone: "81122334455"},
		{LastName: "Смирнов", FirstName: "Алексей", MiddleName: "Сергеевич", Phone: "85566778899"},
		{LastName: "Кузнецов", FirstName: "Дмитрий", MiddleName: "Анатольевич", Phone: "89988776655"},
	}

	parts := []models.Part{
		{Name: "Шина", Description: "Летняя шина", Price: 5000.00},
		{Name: "Аккумулятор", Description: "Аккумулятор 60 Ач", Price: 8000.00},
		{Name: "Фильтр масляный", Description: "Фильтр масляный для автомобиля", Price: 500.00},
		{Name: "Фильтр воздушный", Description: "Фильтр воздушный для автомобиля", Price: 300.00},
		{Name: "Свеча зажигания", Description: "Свеча зажигания для автомобиля", Price: 200.00},
	}

	orders := []models.Order{
		{ClientID: 1, PartID: 1, Quantity: 2, TotalCost: 10000.00},
		{ClientID: 2, PartID: 2, Quantity: 1, TotalCost: 8000.00},
		{ClientID: 3, PartID: 3, Quantity: 4, TotalCost: 2000.00},
		{ClientID: 4, PartID: 4, Quantity: 3, TotalCost: 900.00},
		{ClientID: 5, PartID: 5, Quantity: 5, TotalCost: 1000.00},
	}

	for _, client := range clients {
		mock.ExpectExec("INSERT INTO clients").WithArgs(client.LastName, client.FirstName, client.MiddleName, client.Phone).WillReturnResult(sqlmock.NewResult(1, 1))
	}

	for _, part := range parts {
		mock.ExpectExec("INSERT INTO parts").WithArgs(part.Name, part.Description, part.Price).WillReturnResult(sqlmock.NewResult(1, 1))
	}

	for _, order := range orders {
		mock.ExpectExec("INSERT INTO orders").WithArgs(order.ClientID, order.PartID, order.Quantity, order.TotalCost).WillReturnResult(sqlmock.NewResult(1, 1))
	}

	PopulateDatabase(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
