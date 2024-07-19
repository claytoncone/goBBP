package bbp

import (
	"bytes"
	"fmt"
	"math"
	"testing"
)

func TestLimitZero(t *testing.T) {
	pi := New()
	if len(pi.GetHexDigits(0, 0)) != 0 {
		t.Log("zero limit test failed")
	} else {
		t.Log("zero limit test passed")
	}
}

func TestLimitOne(t *testing.T) {
	pi := New()
	if len(pi.GetHexDigits(0, 1)) != 1 {
		t.Log("one limit test failed")
	} else {
		t.Log("one limit test passed")
	}
}

func TestLimitNine(t *testing.T) {
	pi := New()
	if len(pi.GetHexDigits(0, 9)) != 9 {
		t.Log("nine limit test failed")
	} else {
		t.Log("nine limit test passed")
	}
}
func TestGetDigits(t *testing.T) {
	pi := New()
	digits := pi.GetHexDigits(0, 5)
	if len(digits) != 5 {
		t.Errorf("Expected 5 digits, got %d", len(digits))
	}
}

func TestGetDigitsNegativeArgs(t *testing.T) {
	pi := New()
	digits := pi.GetHexDigits(-1, -5)
	if len(digits) != 0 {
		t.Errorf("Expected 0 digits for negative arguments, got %d", len(digits))
	}
}

func TestGetDigitsStart0(t *testing.T) {
	pi := New()
	expected := []byte{2, 4, 3, 15, 6, 10, 8, 8, 8, 5, 10, 3, 0, 8, 13, 3, 1, 3, 1, 9}

	digits := pi.GetHexDigits(0, len(expected))
	if !bytes.Equal(digits, expected) {
		t.Errorf("Expected %v, got %v", expected, digits)
	}
}
func TestGetDigitsStart8(t *testing.T) {
	pi := New()
	expected := []byte{8, 5, 10, 3, 0, 8, 13, 3, 1, 3, 1, 9}

	digits := pi.GetHexDigits(8, len(expected))
	if !bytes.Equal(digits, expected) {
		t.Errorf("Expected %v, got %v", expected, digits)
	}
}
func TestGetDigits_ExceedsClimit(t *testing.T) {
	p := &PiDigits{}

	start := CLIMIT - 10
	num := 20

	digits := p.GetHexDigits(start, num)

	if len(digits) != 0 {
		t.Errorf("Expected 0 digits, got %d", len(digits))
	}
}

func TestGet4Value(t *testing.T) {
	pi := New()
	digits := pi.GetHexDigits(0, 4)
	if len(digits) != 4 {
		t.Fatalf("Expected 4 digits, got %d", len(digits))
	}
	var value, ex float64
	ex = 16
	for i := 0; i < len(digits); i++ {
		v := float64(int8(digits[i]))
		value = value + v/ex
		ex = ex * 16
	}
	var iValue int16 = int16(value * 10000)
	var expected int16 = 1415
	if iValue == expected {
		t.Logf("TestGet4Value passed - %d", iValue)
	} else {
		t.Errorf("TestGet4Value failed - expected %d, got %d", expected, iValue)
	}
}

func TestValues(t *testing.T) {
	var tests = []struct {
		name   string
		num    int
		result int64
	}{
		{"PI digits", 3, 141},
		{"PI digits", 4, 1415},
		{"PI digits", 5, 14159},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %s", tt.num, tt.name), func(t *testing.T) {
			pi := New()
			digits := pi.GetHexDigits(0, tt.num)
			if len(digits) != tt.num {
				t.Errorf("Expected %d digits, got %d", tt.num, len(digits))
			} else {
				var value, ex float64
				value = 0
				ex = 16
				for i := 0; i < len(digits); i++ {
					v := float64(int8(digits[i]))
					value += v / ex
					ex = ex * 16
				}
				var ivalue int64 = int64(value * math.Pow(10, float64(tt.num)))
				if ivalue != tt.result {
					t.Errorf("Expected %d, got %d", tt.result, ivalue)
				} else {
					t.Logf("%d %s passed - %d", tt.num, tt.name, ivalue)
				}
			}
		})
	}
}
