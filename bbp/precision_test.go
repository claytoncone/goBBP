package bbp // or package bbp if split

import (
	"math"
	"testing"
)

const knownHex = "243f6a8885a308d313198a2e03707344a4093822299f31d0082efa98ec4e6c8949604c2bf" // first 80 known hex digits (safe)

// reconstructFrac builds float64 x = 0.d1 d2 d3 ... dn in hex (fractional part)
func reconstructFrac(hexDigits string) float64 {
	if len(hexDigits) > 15 { // float64 limit
		hexDigits = hexDigits[:15]
	}
	x := 0.0
	power := 1.0 / 16.0
	for _, c := range hexDigits {
		d := hexCharToInt(byte(c))
		x += float64(d) * power
		power /= 16.0
	}
	return x
}

func hexCharToInt(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	return 10 + int(c-'a') // assume lowercase
}

// TestIHexMaxStep finds the largest STEP where chained ihex recovers known digits accurately
func TestIHexMaxStep(t *testing.T) {
	pi := New()
	// Use enough known digits to test beyond float64 limit
	fullKnown := knownHex[:60]

	for step := 12; step >= 1; step-- { // backward from large to small, as you did
		// Reconstruct x with all known digits (float64 will lose tail precision)
		x := reconstructFrac(fullKnown)

		extracted := []byte{}
		currentX := x
		remaining := len(fullKnown)

		for remaining > 0 {
			thisStep := step
			if thisStep > remaining {
				thisStep = remaining
			}

			// Call ihex on currentX
			digits := pi.ihex(currentX, thisStep) // assume returns []byte of hex digits
			extracted = append(extracted, digits...)

			// Update currentX to the remaining fractional part after extracting thisStep digits
			// This simulates your GetHexDigits/genHex update
			// Since ihex extracts by *16 and subtracts integer, the remaining frac is x - floor(x) after shifts
			// But since ihex doesn't return updated x, we simulate by dividing out the extracted digits
			for range digits {
				currentX = currentX * 16
				currentX -= math.Floor(currentX)
			}

			remaining -= thisStep
		}

		got := string(extracted[:len(fullKnown)])
		matchLen := 0
		for i := range got {
			if i >= len(fullKnown) || got[i] != fullKnown[i] {
				break
			}
			matchLen++
		}

		t.Logf("STEP=%d: accurate hex digits = %d (got %s...)", step, matchLen, got[:min(30, matchLen)])

		if step <= 6 && matchLen < 12 {
			t.Errorf("Even small STEP=%d failed basic accuracy (only %d correct)", step, matchLen)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
