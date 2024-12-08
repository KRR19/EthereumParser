package hex

import (
	"testing"
)

func TestToDec_Success(t *testing.T) {
	tests := []struct {
		hex      string
		expected int
	}{
		{"0x1", 1},
		{"0xa", 10},
		{"0x10", 16},
		{"0xff", 255},
	}

	for _, tt := range tests {
		t.Run(tt.hex, func(t *testing.T) {
			result, err := ToDec(tt.hex)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if result != tt.expected {
				t.Fatalf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestToDec_Failure(t *testing.T) {
	tests := []string{
		"0xg",
		"0x",
		"xyz",
	}

	for _, hex := range tests {
		t.Run(hex, func(t *testing.T) {
			_, err := ToDec(hex)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
		})
	}
}
