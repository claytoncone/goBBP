package bbp_test

import (
	"fmt"
	"testing"

	"goBBP/internal/bbp"
)

const (
	testPosition = 10000
	digitCount   = 16 // Number of hex digits to compute per run
	minOverlap   = 12 // Minimum digits that must match (allowing trailing variance)
)

// bytesToHexString converts raw bytes (0-15) to hex character string for display
func bytesToHexString(bytes []byte) string {
	const hexChars = "0123456789ABCDEF"
	result := make([]byte, len(bytes))
	for i, b := range bytes {
		result[i] = hexChars[b&0x0F]
	}
	return string(result)
}

// bytesEqual compares two byte slices
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBBP_OverlapValidation_DMinus1(t *testing.T) {
	pi := bbp.New()

	// Compute digits starting at position d
	digitsAtD := pi.GetHexDigits(testPosition, digitCount)

	// Compute digits starting at position d-1
	digitsAtDMinus1 := pi.GetHexDigits(testPosition-1, digitCount)

	t.Logf("Position %d:   %s", testPosition, bytesToHexString(digitsAtD))
	t.Logf("Position %d: %s", testPosition-1, bytesToHexString(digitsAtDMinus1))

	// The digits at d should match digits at d-1 shifted by 1
	overlap := digitsAtDMinus1[1 : minOverlap+1]
	expected := digitsAtD[0:minOverlap]

	if !bytesEqual(overlap, expected) {
		t.Errorf("Overlap mismatch (d-1 validation)\n"+
			"  Position %d[1:%d]: %s\n"+
			"  Position %d[0:%d]: %s",
			testPosition-1, minOverlap+1, bytesToHexString(overlap),
			testPosition, minOverlap, bytesToHexString(expected))
	} else {
		t.Logf("✓ Overlap validated: %d digits match", minOverlap)
	}
}

func TestBBP_OverlapValidation_DPlus1(t *testing.T) {
	pi := bbp.New()

	digitsAtD := pi.GetHexDigits(testPosition, digitCount)
	digitsAtDPlus1 := pi.GetHexDigits(testPosition+1, digitCount)

	t.Logf("Position %d: %s", testPosition, bytesToHexString(digitsAtD))
	t.Logf("Position %d:  %s", testPosition+1, bytesToHexString(digitsAtDPlus1))

	// The digits at d+1 should match digits at d shifted by 1
	overlap := digitsAtD[1 : minOverlap+1]
	expected := digitsAtDPlus1[0:minOverlap]

	if !bytesEqual(overlap, expected) {
		t.Errorf("Overlap mismatch (d+1 validation)\n"+
			"  Position %d[1:%d]: %s\n"+
			"  Position %d[0:%d]: %s",
			testPosition, minOverlap+1, bytesToHexString(overlap),
			testPosition+1, minOverlap, bytesToHexString(expected))
	} else {
		t.Logf("✓ Overlap validated: %d digits match", minOverlap)
	}
}

func TestBBP_OverlapValidation_BothDirections(t *testing.T) {
	pi := bbp.New()

	digitsAtD := pi.GetHexDigits(testPosition, digitCount)
	digitsAtDMinus1 := pi.GetHexDigits(testPosition-1, digitCount)
	digitsAtDPlus1 := pi.GetHexDigits(testPosition+1, digitCount)

	t.Logf("Position %d: %s", testPosition-1, bytesToHexString(digitsAtDMinus1))
	t.Logf("Position %d:  %s", testPosition, bytesToHexString(digitsAtD))
	t.Logf("Position %d:   %s", testPosition+1, bytesToHexString(digitsAtDPlus1))

	// Check d-1 overlap
	if !bytesEqual(digitsAtDMinus1[1:minOverlap+1], digitsAtD[0:minOverlap]) {
		t.Error("d-1 overlap validation failed")
	}

	// Check d+1 overlap
	if !bytesEqual(digitsAtD[1:minOverlap+1], digitsAtDPlus1[0:minOverlap]) {
		t.Error("d+1 overlap validation failed")
	}

	// Check transitive property: d-1 shifted by 2 should match d+1
	if !bytesEqual(digitsAtDMinus1[2:minOverlap], digitsAtDPlus1[0:minOverlap-2]) {
		t.Error("Transitive validation (d-1 to d+1) failed")
	}

	t.Log("✓ All overlap validations passed")
}

func TestBBP_OverlapValidation_MultiplePositions(t *testing.T) {
	positions := []int{1000, 5000, 10000, 50000, 100000}

	for _, pos := range positions {
		pos := pos // capture range variable
		t.Run(fmt.Sprintf("Position_%d", pos), func(t *testing.T) {
			pi := bbp.New()

			digitsAtD := pi.GetHexDigits(pos, digitCount)
			digitsAtDMinus1 := pi.GetHexDigits(pos-1, digitCount)

			overlap := digitsAtDMinus1[1 : minOverlap+1]
			expected := digitsAtD[0:minOverlap]

			if !bytesEqual(overlap, expected) {
				t.Errorf("Position %d: overlap mismatch\n"+
					"  Got:      %s\n"+
					"  Expected: %s",
					pos, bytesToHexString(overlap), bytesToHexString(expected))
			} else {
				t.Logf("Position %d: ✓ %d digits match (%s...)",
					pos, minOverlap, bytesToHexString(expected[:6]))
			}
		})
	}
}

// TestBBP_DecimalDigitsInRange verifies decimal output is 0-9
func TestBBP_DecimalDigitsInRange(t *testing.T) {
	pi := bbp.New()

	positions := []int{0, 1000, 10000, 100000}

	for _, pos := range positions {
		digits := pi.GetDecimalValues(pos, 20)

		for i, d := range digits {
			if d > 9 {
				t.Errorf("Position %d, digit %d: got %d, want 0-9", pos, i, d)
			}
		}
	}

	t.Log("✓ All decimal digits in valid range 0-9")
}

// TestBBP_KnownDecimalDigits validates against known Pi decimal digits
// Pi = 3.14159265358979323846264338327950288419716939937510...
func TestBBP_KnownDecimalDigits(t *testing.T) {
	pi := bbp.New()

	// Known decimal digits of Pi after the decimal point (position 0 = '1')
	// 3.14159265358979323846...
	knownDecimal := []byte{1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2}

	computed := pi.GetDecimalValues(0, 16)

	t.Logf("Known:    %v", knownDecimal)
	t.Logf("Computed: %v", computed)

	// Count matching leading digits
	matchCount := 0
	for i := 0; i < len(knownDecimal) && i < len(computed); i++ {
		if knownDecimal[i] == computed[i] {
			matchCount++
		} else {
			t.Logf("First mismatch at position %d: expected %d, got %d",
				i, knownDecimal[i], computed[i])
			break
		}
	}

	// Allow some tolerance due to floating point precision
	if matchCount < 10 {
		t.Errorf("Only %d leading digits match known Pi values (expected >= 10)", matchCount)
	} else {
		t.Logf("✓ %d leading digits match known Pi decimal values", matchCount)
	}
}

// TestBBP_DecimalConsistency verifies same input produces same output
func TestBBP_DecimalConsistency(t *testing.T) {
	pi1 := bbp.New()
	pi2 := bbp.New()

	digits1 := pi1.GetDecimalValues(10000, 20)
	digits2 := pi2.GetDecimalValues(10000, 20)

	if !bytesEqual(digits1, digits2) {
		t.Errorf("Inconsistent results:\n  Run 1: %v\n  Run 2: %v", digits1, digits2)
	} else {
		t.Logf("✓ Decimal computation is consistent: %v", digits1)
	}
}

// TestBBP_KnownDigits validates against known Pi hex digits
// Position 0 of Pi in hex is 243F6A8885A308D3...
func TestBBP_KnownDigits(t *testing.T) {
	pi := bbp.New()

	// Known hex digits of Pi starting at position 0 (after "3.")
	// Pi = 3.243F6A8885A308D313198A2E03707344A4093822299F31D0...
	knownHex := []byte{2, 4, 3, 15, 6, 10, 8, 8, 8, 5, 10, 3, 0, 8, 13, 3}

	computed := pi.GetHexDigits(0, 16)

	t.Logf("Known:    %s", bytesToHexString(knownHex))
	t.Logf("Computed: %s", bytesToHexString(computed))

	// Allow some trailing digit variance due to floating point
	matchCount := 0
	for i := 0; i < len(knownHex) && i < len(computed); i++ {
		if knownHex[i] == computed[i] {
			matchCount++
		} else {
			break
		}
	}

	if matchCount < 12 {
		t.Errorf("Only %d leading digits match known Pi values (expected >= 12)", matchCount)
	} else {
		t.Logf("✓ %d leading digits match known Pi values", matchCount)
	}
}
