package main

import (
	"autoservice/internal/db"
	"autoservice/internal/handlers"
	"fmt"
	"log"
)

func main() {
	// Открываем или создаем базу данных
	dbConn, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Создаем таблицы, если они не существуют
	db.CreateTables(dbConn)

	// Основной цикл программы
	for {
		fmt.Println("\nВыберите действие:")
		fmt.Println("1. Добавить клиента")
		fmt.Println("2. Добавить деталь")
		fmt.Println("3. Добавить заказ")
		fmt.Println("4. Удалить клиента")
		fmt.Println("5. Удалить деталь")
		fmt.Println("6. Удалить заказ")
		fmt.Println("7. Вывести всех клиентов")
		fmt.Println("8. Вывести все детали")
		fmt.Println("9. Вывести все заказы")
		fmt.Println("10. Очистить базу данных")
		fmt.Println("11. Заполнить базу данных шаблонными данными")
		fmt.Println("12. Вывести все данные")
		fmt.Println("0. Выход")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			handlers.AddClientMenu(dbConn)
		case 2:
			handlers.AddPartMenu(dbConn)
		case 3:
			handlers.AddOrderMenu(dbConn)
		case 4:
			handlers.DeleteClientMenu(dbConn)
		case 5:
			handlers.DeletePartMenu(dbConn)
		case 6:
			handlers.DeleteOrderMenu(dbConn)
		case 7:
			handlers.PrintClients(dbConn)
		case 8:
			handlers.PrintParts(dbConn)
		case 9:
			handlers.PrintOrders(dbConn)
		case 10:
			db.ClearDatabase(dbConn)
		case 11:
			db.PopulateDatabase(dbConn)
		case 12:
			handlers.PrintClients(dbConn)
			handlers.PrintParts(dbConn)
			handlers.PrintOrders(dbConn)
		case 0:
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
