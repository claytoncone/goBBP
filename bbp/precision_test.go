package bbp

import (
	"fmt"
	"math"
	"testing"
)

const knownHexFrac = "243f6a8885a308d313198a2e03707344a4093822299f31d0082efa98ec4e6c8949604c2bf" // first 80 known hex digits (safe reference)

// helper to create float64 with known hex fractional part (up to ~15 accurate digits)
func makeFracFromHex(hexStr string) float64 {
	if len(hexStr) > 15 {
		hexStr = hexStr[:15] // float64 limit
	}
	x := 0.0
	for _, c := range hexStr {
		d := hexToInt(byte(c))
		x = x*16 + float64(d)
	}
	return x / math.Pow(16, float64(len(hexStr)))
}

func hexToInt(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	return 10 + int(c-'a')
}

// TestIHexChainedExtraction tests chained ihex calls with varying STEP
func TestIHexChainedExtraction(t *testing.T) {

	pi := New()

	// Use known digits (more than float64 can hold accurately)
	fullKnown := knownHexFrac[:60] // plenty

	for step := 6; step <= 12; step++ {
		t.Run(fmt.Sprintf("STEP=%d", step), func(t *testing.T) {
			// Start with x containing ~15 known digits
			x := makeFracFromHex(fullKnown[:15])

			extracted := []byte{}
			currentX := x
			remaining := 50 // target total digits

			for remaining > 0 {
				thisStep := step
				if thisStep > remaining {
					thisStep = remaining
				}
				digits := pi.ihex(currentX, thisStep)
				extracted = append(extracted, digits...)

				// Update currentX for next batch (simulate chained frac update)
				// ihex typically returns digits and leaves updated frac in x or returns it
				// Adjust if your ihex modifies x in place or returns new x
				// Example if ihex returns updated frac:
				// currentX = updatedFrac
				// Or if in place:
				currentX = currentX - math.Floor(currentX) // crude — replace with your exact update

				remaining -= thisStep
			}

			got := string(extracted[:50])
			expectedPrefix := fullKnown[:50]
			matchLen := 0
			for i := range got {
				if i >= len(expectedPrefix) || got[i] != expectedPrefix[i] {
					break
				}
				matchLen++
			}

			t.Logf("STEP=%d: accurate digits = %d (got prefix %s)", step, matchLen, got[:matchLen])

			if matchLen < 10 {
				t.Errorf("STEP=%d too large for basic accuracy (only %d correct digits)", step, matchLen)
			}
		})
	}
}
