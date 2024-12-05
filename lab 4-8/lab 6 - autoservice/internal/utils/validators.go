package utils

import (
	"regexp"
	"unicode"
)

// IsValidName проверяет, что строка начинается с заглавной буквы
func IsValidName(name string) bool {
	if len(name) == 0 {
		return false
	}
	firstChar := []rune(name)[0]
	return unicode.IsUpper(firstChar)
}

// IsValidPhone проверяет, что номер телефона начинается с цифры "8" и состоит только из цифр
func IsValidPhone(phone string) bool {
	re := regexp.MustCompile(`^8\d{10}$`)
	return re.MatchString(phone)
}

// IsValidPrice проверяет, что цена является положительным числом
func IsValidPrice(price float64) bool {
	return price > 0
}

// IsValidQuantity проверяет, что количество является положительным целым числом
func IsValidQuantity(quantity int) bool {
	return quantity > 0
}

// IsValidID проверяет, что ID является положительным целым числом
func IsValidID(id int) bool {
	return id > 0
}
