// internal/utils/validators_test.go
package utils

import (
	"testing"
)

func TestIsValidName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", false},
		{"Valid name", "John", true},
		{"Lowercase first letter", "john", false},
		{"Digit first letter", "1John", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidName(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsValidPhone(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid phone", "81234567890", true},
		{"Invalid phone (starts with 7)", "71234567890", false},
		{"Invalid phone (too short)", "8123456789", false},
		{"Invalid phone (contains letters)", "8123456789a", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPhone(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsValidPrice(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected bool
	}{
		{"Positive price", 100.00, true},
		{"Zero price", 0.00, false},
		{"Negative price", -100.00, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPrice(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsValidQuantity(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Positive quantity", 10, true},
		{"Zero quantity", 0, false},
		{"Negative quantity", -10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidQuantity(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsValidID(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Positive ID", 1, true},
		{"Zero ID", 0, false},
		{"Negative ID", -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidID(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
