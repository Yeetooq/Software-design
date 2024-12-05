package handlers

import (
	"autoservice/internal/models"
	"autoservice/internal/utils"
	"database/sql"
	"fmt"
	"log"
)

func AddOrderMenu(db *sql.DB) {
	var clientID, partID, quantity int

	for {
		fmt.Println("Введите ID клиента:")
		fmt.Scanln(&clientID)
		if !utils.IsValidID(clientID) {
			fmt.Println("ID клиента должно быть положительным целым числом. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите ID детали:")
		fmt.Scanln(&partID)
		if !utils.IsValidID(partID) {
			fmt.Println("ID детали должно быть положительным целым числом. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите количество:")
		fmt.Scanln(&quantity)
		if !utils.IsValidQuantity(quantity) {
			fmt.Println("Количество деталей должно быть положительным целым числом. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	var totalCost float64
	err := db.QueryRow("SELECT price FROM parts WHERE id = ?", partID).Scan(&totalCost)
	if err != nil {
		log.Println(err)
		return
	}
	totalCost *= float64(quantity)

	order := models.Order{ClientID: clientID, PartID: partID, Quantity: quantity, TotalCost: totalCost}
	_, err = addOrder(db, order)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Заказ успешно добавлен.")
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

func DeleteOrderMenu(db *sql.DB) {
	var id int
	fmt.Println("Введите ID заказа для удаления:")
	fmt.Scanln(&id)

	err := deleteOrder(db, id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Заказ успешно удален.")
}

func deleteOrder(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM orders WHERE id = ?", id)
	return err
}

func PrintOrders(db *sql.DB) {
	rows, err := db.Query("SELECT id, client_id, part_id, quantity, total_cost FROM orders")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("Заказы:")
	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.ClientID, &order.PartID, &order.Quantity, &order.TotalCost)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("ID: %d, Клиент ID: %d, Деталь ID: %d, Количество: %d, Общая стоимость: %.2f\n", order.ID, order.ClientID, order.PartID, order.Quantity, order.TotalCost)
	}
}
