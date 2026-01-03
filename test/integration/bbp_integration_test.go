//go:build integration
// +build integration

package integration

import (
	"fmt"
	"testing"

	"goBBP/internal/bbp"
	"goBBP/internal/convert"
)

const (
	digitCount = 16 // Number of hex digits to compute per run
	minOverlap = 12 // Minimum digits that must match
)

// --- Overlap Validation Tests ---
// These tests verify the BBP implementation using the overlap method:
// "The result for position d can be easily checked by computing the result
// for position d-1. If the results perfectly overlap with an offset of 1,
// this is a strong indication that the code is working properly."

func TestBBP_OverlapValidation_DMinus1(t *testing.T) {
	testPosition := 10000
	pi := bbp.New()

	digitsAtD := pi.GetHexDigits(testPosition, digitCount)
	digitsAtDMinus1 := pi.GetHexDigits(testPosition-1, digitCount)

	t.Logf("Position %d:   %s", testPosition, convert.HexBytesToString(digitsAtD))
	t.Logf("Position %d: %s", testPosition-1, convert.HexBytesToString(digitsAtDMinus1))

	overlap := digitsAtDMinus1[1 : minOverlap+1]
	expected := digitsAtD[0:minOverlap]

	if !convert.BytesEqual(overlap, expected) {
		t.Errorf("Overlap mismatch (d-1 validation)\n"+
			"  Position %d[1:%d]: %s\n"+
			"  Position %d[0:%d]: %s",
			testPosition-1, minOverlap+1, convert.HexBytesToString(overlap),
			testPosition, minOverlap, convert.HexBytesToString(expected))
	} else {
		t.Logf("✓ Overlap validated: %d digits match", minOverlap)
	}
}

func TestBBP_OverlapValidation_DPlus1(t *testing.T) {
	testPosition := 10000
	pi := bbp.New()

	digitsAtD := pi.GetHexDigits(testPosition, digitCount)
	digitsAtDPlus1 := pi.GetHexDigits(testPosition+1, digitCount)

	t.Logf("Position %d: %s", testPosition, convert.HexBytesToString(digitsAtD))
	t.Logf("Position %d:  %s", testPosition+1, convert.HexBytesToString(digitsAtDPlus1))

	overlap := digitsAtD[1 : minOverlap+1]
	expected := digitsAtDPlus1[0:minOverlap]

	if !convert.BytesEqual(overlap, expected) {
		t.Errorf("Overlap mismatch (d+1 validation)\n"+
			"  Position %d[1:%d]: %s\n"+
			"  Position %d[0:%d]: %s",
			testPosition, minOverlap+1, convert.HexBytesToString(overlap),
			testPosition+1, minOverlap, convert.HexBytesToString(expected))
	} else {
		t.Logf("✓ Overlap validated: %d digits match", minOverlap)
	}
}

func TestBBP_OverlapValidation_BothDirections(t *testing.T) {
	testPosition := 10000
	pi := bbp.New()

	digitsAtD := pi.GetHexDigits(testPosition, digitCount)
	digitsAtDMinus1 := pi.GetHexDigits(testPosition-1, digitCount)
	digitsAtDPlus1 := pi.GetHexDigits(testPosition+1, digitCount)

	t.Logf("Position %d: %s", testPosition-1, convert.HexBytesToString(digitsAtDMinus1))
	t.Logf("Position %d:  %s", testPosition, convert.HexBytesToString(digitsAtD))
	t.Logf("Position %d:   %s", testPosition+1, convert.HexBytesToString(digitsAtDPlus1))

	// Check d-1 overlap
	if !convert.BytesEqual(digitsAtDMinus1[1:minOverlap+1], digitsAtD[0:minOverlap]) {
		t.Error("d-1 overlap validation failed")
	}

	// Check d+1 overlap
	if !convert.BytesEqual(digitsAtD[1:minOverlap+1], digitsAtDPlus1[0:minOverlap]) {
		t.Error("d+1 overlap validation failed")
	}

	// Check transitive property: d-1 shifted by 2 should match d+1
	if !convert.BytesEqual(digitsAtDMinus1[2:minOverlap], digitsAtDPlus1[0:minOverlap-2]) {
		t.Error("Transitive validation (d-1 to d+1) failed")
	}

	t.Log("✓ All overlap validations passed")
}

func TestBBP_OverlapValidation_MultiplePositions(t *testing.T) {
	positions := []int{1000, 5000, 10000, 50000, 100000}

	for _, pos := range positions {
		pos := pos
		t.Run(fmt.Sprintf("Position_%d", pos), func(t *testing.T) {
			pi := bbp.New()

			digitsAtD := pi.GetHexDigits(pos, digitCount)
			digitsAtDMinus1 := pi.GetHexDigits(pos-1, digitCount)

			overlap := digitsAtDMinus1[1 : minOverlap+1]
			expected := digitsAtD[0:minOverlap]

			if !convert.BytesEqual(overlap, expected) {
				t.Errorf("Position %d: overlap mismatch\n"+
					"  Got:      %s\n"+
					"  Expected: %s",
					pos, convert.HexBytesToString(overlap), convert.HexBytesToString(expected))
			} else {
				t.Logf("Position %d: ✓ %d digits match (%s...)",
					pos, minOverlap, convert.HexBytesToString(expected[:6]))
			}
		})
	}
}

// --- Known Value Tests ---

func TestBBP_KnownHexDigits(t *testing.T) {
	pi := bbp.New()

	// Known hex digits of Pi starting at position 0 (after "3.")
	// Pi = 3.243F6A8885A308D313198A2E03707344A4093822299F31D0...
	knownHex := []byte{2, 4, 3, 15, 6, 10, 8, 8, 8, 5, 10, 3, 0, 8, 13, 3}
	computed := pi.GetHexDigits(0, 16)

	t.Logf("Known:    %s", convert.HexBytesToString(knownHex))
	t.Logf("Computed: %s", convert.HexBytesToString(computed))

	matchCount := 0
	for i := 0; i < len(knownHex) && i < len(computed); i++ {
		if knownHex[i] == computed[i] {
			matchCount++
		} else {
			t.Logf("First mismatch at position %d: expected %X, got %X",
				i, knownHex[i], computed[i])
			break
		}
	}

	if matchCount < 12 {
		t.Errorf("Only %d leading digits match (expected >= 12)", matchCount)
	} else {
		t.Logf("✓ %d leading digits match known Pi hex values", matchCount)
	}
}

// --- Consistency Tests ---

func TestBBP_HexConsistency(t *testing.T) {
	pi1 := bbp.New()
	pi2 := bbp.New()

	digits1 := pi1.GetHexDigits(10000, 20)
	digits2 := pi2.GetHexDigits(10000, 20)

	if !convert.BytesEqual(digits1, digits2) {
		t.Errorf("Inconsistent results:\n  Run 1: %s\n  Run 2: %s",
			convert.HexBytesToString(digits1), convert.HexBytesToString(digits2))
	} else {
		t.Logf("✓ Hex computation is consistent: %s", convert.HexBytesToString(digits1))
	}
}

// --- Edge Case Tests ---

func TestBBP_InvalidInputs(t *testing.T) {
	pi := bbp.New()

	tests := []struct {
		name  string
		start int
		num   int
	}{
		{"negative start", -1, 10},
		{"zero count", 0, 0},
		{"negative count", 0, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pi.GetHexDigits(tt.start, tt.num)
			if len(result) != 0 {
				t.Errorf("Expected empty result, got %d bytes", len(result))
			}
		})
	}

	t.Log("✓ Invalid inputs handled correctly")
}

// --- String Output Tests ---

func TestBBP_GetHexString(t *testing.T) {
	pi := bbp.New()

	result := pi.GetHexString(0, 16)
	expected := "243F6A8885A308D3"

	t.Logf("GetHexString: %s", result)

	matchLen := 12
	if len(result) < matchLen {
		t.Fatalf("Result too short: got %d chars", len(result))
	}

	if result[:matchLen] != expected[:matchLen] {
		t.Errorf("Expected prefix %s, got %s", expected[:matchLen], result[:matchLen])
	} else {
		t.Logf("✓ Hex string output correct")
	}
}

func TestBBP_GetHexValues(t *testing.T) {
	pi := bbp.New()

	result := pi.GetHexValues(0, 16)
	expected := []byte("243F6A8885A308D3")

	t.Logf("GetHexValues: %s", string(result))

	matchLen := 12
	if !convert.BytesEqual(result[:matchLen], expected[:matchLen]) {
		t.Errorf("Expected prefix %s, got %s", string(expected[:matchLen]), string(result[:matchLen]))
	} else {
		t.Logf("✓ Hex values output correct")
	}
}
