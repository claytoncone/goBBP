package convert

import (
	"testing"
)

func TestHexBytesToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"basic", []byte{2, 4, 3, 15, 6, 10}, "243F6A"},
		{"all digits", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, "0123456789ABCDEF"},
		{"empty", []byte{}, ""},
		{"nil", nil, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HexBytesToString(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestHexBytesToLowerString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"basic", []byte{2, 4, 3, 15, 6, 10}, "243f6a"},
		{"all digits", []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, "0123456789abcdef"},
		{"empty", []byte{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HexBytesToLowerString(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestHexBytesToAlpha(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{"basic", []byte{2, 4, 3, 15, 6, 10}, []byte{'2', '4', '3', 'F', '6', 'A'}},
		{"empty", []byte{}, nil},
		{"nil", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HexBytesToAlpha(tt.input)
			if !BytesEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestDecBytesToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"basic", []byte{1, 4, 1, 5, 9}, "14159"},
		{"with zero", []byte{3, 0, 4, 1}, "3041"},
		{"empty", []byte{}, ""},
		{"nil", nil, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DecBytesToString(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestBytesEqual(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []byte
		expected bool
	}{
		{"equal", []byte{1, 2, 3}, []byte{1, 2, 3}, true},
		{"not equal", []byte{1, 2, 3}, []byte{1, 2, 4}, false},
		{"different length", []byte{1, 2}, []byte{1, 2, 3}, false},
		{"both nil", nil, nil, true},
		{"both empty", []byte{}, []byte{}, true},
		{"nil vs empty", nil, []byte{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BytesEqual(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// --- Benchmarks ---

func BenchmarkHexBytesToString(b *testing.B) {
	hex := make([]byte, 100)
	for i := range hex {
		hex[i] = byte(i % 16)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HexBytesToString(hex)
	}
}

func BenchmarkBytesEqual(b *testing.B) {
	a := make([]byte, 100)
	c := make([]byte, 100)
	for i := range a {
		a[i] = byte(i)
		c[i] = byte(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BytesEqual(a, c)
	}
}
