package handlers

import (
	"autoservice/internal/models"
	"autoservice/internal/utils"
	"database/sql"
	"fmt"
	"log"
)

func AddPartMenu(db *sql.DB) {
	var name, description string
	var price float64

	for {
		fmt.Println("Введите название детали:")
		fmt.Scanln(&name)
		if !utils.IsValidName(name) {
			fmt.Println("Название детали должно начинаться с заглавной буквы. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите описание детали:")
		fmt.Scanln(&description)
		if !utils.IsValidName(description) {
			fmt.Println("Описание детали должно начинаться с заглавной буквы. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	for {
		fmt.Println("Введите цену детали:")
		fmt.Scanln(&price)
		if !utils.IsValidPrice(price) {
			fmt.Println("Цена детали должна быть положительным числом. Пожалуйста, введите данные заново.")
			continue
		}
		break
	}

	part := models.Part{Name: name, Description: description, Price: price}
	_, err := addPart(db, part)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Деталь успешно добавлена.")
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

func DeletePartMenu(db *sql.DB) {
	var id int
	fmt.Println("Введите ID детали для удаления:")
	fmt.Scanln(&id)

	err := deletePart(db, id)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("Деталь успешно удалена.")
}

func deletePart(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM parts WHERE id = ?", id)
	return err
}

func PrintParts(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, description, price FROM parts")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("Детали:")
	for rows.Next() {
		var part models.Part
		err := rows.Scan(&part.ID, &part.Name, &part.Description, &part.Price)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("ID: %d, Название: %s, Описание: %s, Цена: %.2f\n", part.ID, part.Name, part.Description, part.Price)
	}
}
