// cmd/main_test.go
package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainFunction(t *testing.T) {
	// Сохраняем оригинальные значения переменных окружения
	originalArgs := os.Args
	originalStdin := os.Stdin
	originalStdout := os.Stdout
	defer func() {
		os.Args = originalArgs
		os.Stdin = originalStdin
		os.Stdout = originalStdout
	}()

	// Тестовый случай 1: Проверка работы программы без аргументов
	os.Args = []string{"autoservice"}
	capturedOutput := captureOutput(main)
	if !strings.Contains(capturedOutput, "Выберите действие:") {
		t.Errorf("expected 'Выберите действие:', got %s", capturedOutput)
	}

	// Тестовый случай 2: Проверка работы программы с выбором "0" (выход)
	os.Args = []string{"autoservice"}
	input := "0\n"
	r, w, _ := os.Pipe()
	_, err := io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if capturedOutput != "" {
		t.Errorf("expected no output, got %s", capturedOutput)
	}

	// Тестовый случай 3: Проверка работы программы с выбором "1" (добавить клиента)
	os.Args = []string{"autoservice"}
	input = "1\nDoe\nJohn\nSmith\n81234567890\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Клиент успешно добавлен.") {
		t.Errorf("expected 'Клиент успешно добавлен.', got %s", capturedOutput)
	}

	// Тестовый случай 4: Проверка работы программы с выбором "2" (добавить деталь)
	os.Args = []string{"autoservice"}
	input = "2\nWheel\nCar wheel\n100.00\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Деталь успешно добавлена.") {
		t.Errorf("expected 'Деталь успешно добавлена.', got %s", capturedOutput)
	}

	// Тестовый случай 5: Проверка работы программы с выбором "3" (добавить заказ)
	os.Args = []string{"autoservice"}
	input = "3\n1\n1\n2\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Заказ успешно добавлен.") {
		t.Errorf("expected 'Заказ успешно добавлен.', got %s", capturedOutput)
	}

	// Тестовый случай 6: Проверка работы программы с выбором "4" (удалить клиента)
	os.Args = []string{"autoservice"}
	input = "4\n1\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Клиент успешно удален.") {
		t.Errorf("expected 'Клиент успешно удален.', got %s", capturedOutput)
	}

	// Тестовый случай 7: Проверка работы программы с выбором "5" (удалить деталь)
	os.Args = []string{"autoservice"}
	input = "5\n1\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Деталь успешно удалена.") {
		t.Errorf("expected 'Деталь успешно удалена.', got %s", capturedOutput)
	}

	// Тестовый случай 8: Проверка работы программы с выбором "6" (удалить заказ)
	os.Args = []string{"autoservice"}
	input = "6\n1\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Заказ успешно удален.") {
		t.Errorf("expected 'Заказ успешно удален.', got %s", capturedOutput)
	}

	// Тестовый случай 9: Проверка работы программы с выбором "7" (вывести всех клиентов)
	os.Args = []string{"autoservice"}
	input = "7\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Клиенты:") {
		t.Errorf("expected 'Клиенты:', got %s", capturedOutput)
	}

	// Тестовый случай 10: Проверка работы программы с выбором "8" (вывести все детали)
	os.Args = []string{"autoservice"}
	input = "8\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Детали:") {
		t.Errorf("expected 'Детали:', got %s", capturedOutput)
	}

	// Тестовый случай 11: Проверка работы программы с выбором "9" (вывести все заказы)
	os.Args = []string{"autoservice"}
	input = "9\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Заказы:") {
		t.Errorf("expected 'Заказы:', got %s", capturedOutput)
	}

	// Тестовый случай 12: Проверка работы программы с выбором "10" (очистить базу данных)
	os.Args = []string{"autoservice"}
	input = "10\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "База данных успешно очищена.") {
		t.Errorf("expected 'База данных успешно очищена.', got %s", capturedOutput)
	}

	// Тестовый случай 13: Проверка работы программы с выбором "11" (заполнить базу данных шаблонными данными)
	os.Args = []string{"autoservice"}
	input = "11\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "База данных успешно заполнена шаблонными данными.") {
		t.Errorf("expected 'База данных успешно заполнена шаблонными данными.', got %s", capturedOutput)
	}

	// Тестовый случай 14: Проверка работы программы с выбором "12" (вывести все данные)
	os.Args = []string{"autoservice"}
	input = "12\n0\n"
	r, w, _ = os.Pipe()
	_, err = io.WriteString(w, input)
	if err != nil {
		t.Fatalf("failed to write to pipe: %s", err)
	}
	w.Close()
	os.Stdin = r

	capturedOutput = captureOutput(main)
	if !strings.Contains(capturedOutput, "Клиенты:\nДетали:\nЗаказы:") {
		t.Errorf("expected 'Клиенты:\nДетали:\nЗаказы:', got %s", capturedOutput)
	}
}

// captureOutput перехватывает вывод функции
func captureOutput(f func()) string {
	// Перенаправляем стандартный вывод в буфер
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Выполняем функцию
	f()

	// Закрываем буфер и восстанавливаем стандартный вывод
	w.Close()
	os.Stdout = old

	// Читаем данные из буфера
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
