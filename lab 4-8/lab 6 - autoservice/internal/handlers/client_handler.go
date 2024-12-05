package handlers

import (
	"autoservice/internal/models"
	"autoservice/internal/utils"
	"database/sql"
	"fmt"
	"log"
)

func AddClientMenu(db *sql.DB) {
	var lastName, firstName, middleName, phone string

	for {
		fmt.Println("Введите фамилию клиента:")
		fmt.Scanln(&lastName)
		if !utils.IsValidName(lastName) {
			fmt.Println("Фамилия должна начинаться с заглавной буквы. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите имя клиента:")
		fmt.Scanln(&firstName)
		if !utils.IsValidName(firstName) {
			fmt.Println("Имя должно начинаться с заглавной буквы. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите отчество клиента:")
		fmt.Scanln(&middleName)
		if !utils.IsValidName(middleName) {
			fmt.Println("Отчество должно начинаться с заглавной буквы. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите телефон клиента:")
		fmt.Scanln(&phone)
		if !utils.IsValidPhone(phone) {
			fmt.Println("Номер телефона должен начинаться с цифры '8' и состоять только из цифр. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	client := models.Client{LastName: lastName, FirstName: firstName, MiddleName: middleName, Phone: phone}
	_, err := addClient(db, client)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Клиент успешно добавлен.")
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

func DeleteClientMenu(db *sql.DB) {
	var id int
	fmt.Println("Введите ID клиента для удаления:")
	fmt.Scanln(&id)

	err := deleteClient(db, id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Клиент успешно удален.")
}

func deleteClient(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = ?", id)
	return err
}

func PrintClients(db *sql.DB) {
	rows, err := db.Query("SELECT id, last_name, first_name, middle_name, phone FROM clients")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("Клиенты:")
	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client.ID, &client.LastName, &client.FirstName, &client.MiddleName, &client.Phone)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("ID: %d, Фамилия: %s, Имя: %s, Отчество: %s, Телефон: %s\n", client.ID, client.LastName, client.FirstName, client.MiddleName, client.Phone)
	}
}
