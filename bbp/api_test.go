package bbp

import (
	"bytes"
	"testing"
)

func TestLimitZero(t *testing.T) {
	pi := New()
	if len(pi.Get(0, 0)) != 0 {
		t.Log("zero limit test failed")
	} else {
		t.Log("zero limit test passed")
	}
}

func TestLimitOne(t *testing.T) {
	pi := New()
	if len(pi.Get(0, 1)) != 1 {
		t.Log("one limit test failed")
	} else {
		t.Log("one limit test passed")
	}
}

func TestLimitNine(t *testing.T) {
	pi := New()
	if len(pi.Get(0, 9)) != 9 {
		t.Log("nine limit test failed")
	} else {
		t.Log("nine limit test passed")
	}
}
func TestGetDigits(t *testing.T) {
	pi := New()
	digits := pi.Get(0, 5)
	if len(digits) != 5 {
		t.Errorf("Expected 5 digits, got %d", len(digits))
	}
}

func TestGetDigitsNegativeArgs(t *testing.T) {
	pi := New()
	digits := pi.Get(-1, -5)
	if len(digits) != 0 {
		t.Errorf("Expected 0 digits for negative arguments, got %d", len(digits))
	}
}

func TestGetDigitsStart0(t *testing.T) {
	pi := New()
	expected := []byte{2, 4, 3, 15, 6, 10, 8, 8, 8, 5, 10, 3, 0, 8, 13, 3, 1, 3, 1, 9}
	digits := pi.Get(0, len(expected))
	if !bytes.Equal(digits, expected) {
		t.Errorf("Expected %v, got %v", expected, digits)
	}
}
func TestGetDigits_ExceedsClimit(t *testing.T) {
	p := &PiDigits{}

	start := CLIMIT - 10
	num := 20

	digits := p.Get(start, num)

	if len(digits) != 0 {
		t.Errorf("Expected 0 digits, got %d", len(digits))
	}
}
