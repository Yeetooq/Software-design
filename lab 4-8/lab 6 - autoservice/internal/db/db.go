package db

import (
	"autoservice/internal/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Open открывает соединение с базой данных
func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./autoservice.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateTables создает таблицы в базе данных, если они не существуют
func CreateTables(db *sql.DB) {
	createClientsTable := `
	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		last_name TEXT,
		first_name TEXT,
		middle_name TEXT,
		phone TEXT
	);`

	createPartsTable := `
	CREATE TABLE IF NOT EXISTS parts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		description TEXT,
		price REAL
	);`

	createOrdersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER,
		part_id INTEGER,
		quantity INTEGER,
		total_cost REAL,
		FOREIGN KEY (client_id) REFERENCES clients(id),
		FOREIGN KEY (part_id) REFERENCES parts(id)
	);`

	_, err := db.Exec(createClientsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createPartsTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(createOrdersTable)
	if err != nil {
		log.Fatal(err)
	}

	// Обновляем структуру таблицы clients, если необходимо
	updateClientsTable := `
	ALTER TABLE clients ADD COLUMN last_name TEXT;
	ALTER TABLE clients ADD COLUMN first_name TEXT;
	ALTER TABLE clients ADD COLUMN middle_name TEXT;
	`
	_, err = db.Exec(updateClientsTable)
	if err != nil {
		log.Println(err)
	}
}

// ClearDatabase очищает базу данных
func ClearDatabase(db *sql.DB) {
	clearClients := `DELETE FROM clients`
	clearParts := `DELETE FROM parts`
	clearOrders := `DELETE FROM orders`

	_, err := db.Exec(clearClients)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = db.Exec(clearParts)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = db.Exec(clearOrders)
	if err != nil {
		log.Println(err)
		return
	}

	// Сброс автоинкремента для таблиц clients, parts и orders
	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name = 'clients'")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name = 'parts'")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name = 'orders'")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("База данных успешно очищена.")
}

// PopulateDatabase заполняет базу данных шаблонными данными
func PopulateDatabase(db *sql.DB) {
	// Добавляем 5 клиентов
	clients := []models.Client{
		{LastName: "Иванов", FirstName: "Иван", MiddleName: "Иванович", Phone: "81234567890"},
		{LastName: "Петров", FirstName: "Петр", MiddleName: "Петрович", Phone: "80987654321"},
		{LastName: "Сидоров", FirstName: "Сидор", MiddleName: "Сидорович", Phone: "81122334455"},
		{LastName: "Смирнов", FirstName: "Алексей", MiddleName: "Сергеевич", Phone: "85566778899"},
		{LastName: "Кузнецов", FirstName: "Дмитрий", MiddleName: "Анатольевич", Phone: "89988776655"},
	}
	for _, client := range clients {
		_, err := addClient(db, client)
		if err != nil {
			log.Println("Ошибка при добавлении клиента:", err)
		}
	}

	// Добавляем 5 деталей
	parts := []models.Part{
		{Name: "Шина", Description: "Летняя шина", Price: 5000.00},
		{Name: "Аккумулятор", Description: "Аккумулятор 60 Ач", Price: 8000.00},
		{Name: "Фильтр масляный", Description: "Фильтр масляный для автомобиля", Price: 500.00},
		{Name: "Фильтр воздушный", Description: "Фильтр воздушный для автомобиля", Price: 300.00},
		{Name: "Свеча зажигания", Description: "Свеча зажигания для автомобиля", Price: 200.00},
	}
	for _, part := range parts {
		_, err := addPart(db, part)
		if err != nil {
			log.Println("Ошибка при добавлении детали:", err)
		}
	}

	// Добавляем 5 заказов
	orders := []models.Order{
		{ClientID: 1, PartID: 1, Quantity: 2, TotalCost: 10000.00},
		{ClientID: 2, PartID: 2, Quantity: 1, TotalCost: 8000.00},
		{ClientID: 3, PartID: 3, Quantity: 4, TotalCost: 2000.00},
		{ClientID: 4, PartID: 4, Quantity: 3, TotalCost: 900.00},
		{ClientID: 5, PartID: 5, Quantity: 5, TotalCost: 1000.00},
	}
	for _, order := range orders {
		_, err := addOrder(db, order)
		if err != nil {
			log.Println("Ошибка при добавлении заказа:", err)
		}
	}

	fmt.Println("База данных успешно заполнена шаблонными данными.")
}

func addClient(db *sql.DB, client models.Client) (int, error) {
	res, err := db.Exec("INSERT INTO clients (last_name, first_name, middle_name, phone) VALUES (?, ?, ?, ?)", client.LastName, client.FirstName, client.MiddleName, client.Phone)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func addPart(db *sql.DB, part models.Part) (int, error) {
	res, err := db.Exec("INSERT INTO parts (name, description, price) VALUES (?, ?, ?)", part.Name, part.Description, part.Price)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func addOrder(db *sql.DB, order models.Order) (int, error) {
	res, err := db.Exec("INSERT INTO orders (client_id, part_id, quantity, total_cost) VALUES (?, ?, ?, ?)", order.ClientID, order.PartID, order.Quantity, order.TotalCost)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
