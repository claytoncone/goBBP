package bbp

import (
	"encoding/hex"
	"math"
	"strconv"
	"testing"
)

const knownHex = "243f6a8885a308d313198a2e03707344a4093822299f31d0082efa98ec4e6c8949604c2bf" // first 80+ known hex digits

func reconstructFrac(hexDigits string) float64 {
	x := 0.0
	//	reverseHex := Reverse(hexDigits)
	for _, c := range hexDigits {
		d, _ := strconv.ParseInt(string(c), 16, 64)
		x = x*16 + float64(d)
	}
	return x / math.Pow(16, float64(len(hexDigits)))
}
func compressAndConvertToHex(digits []byte) string {
	var compressed []byte
	var i int = 0

	for _, c := range digits {
		if i%2 == 0 {
			compressed = append(compressed, byte(c*16))
		} else {
			compressed[len(compressed)-1] += byte(c)
		}
		i++
	}

	return hex.EncodeToString([]byte(compressed))
}

func TestIHexChainedStepAccuracy(t *testing.T) {
	extra := 2 // STEP + extra for underflow detection
	maxBatches := 10

	for step := 12; step >= 1; step-- {
		t.Run(strconv.Itoa(step), func(t *testing.T) {
			pos := 0 // position in knownHex
			batch := 0
			matchedBatches := 0

			for batch < maxBatches {
				batch++
				start := pos
				end := pos + step + extra
				if end > len(knownHex) {
					end = len(knownHex)
				}
				slice := knownHex[start:end]
				if len(slice) < step {
					break // no more
				}

				x := reconstructFrac(slice)

				pi := PiDigits{} // new instance if needed
				digits := pi.ihex(x, step)

				got := compressAndConvertToHex(digits)
				expected := knownHex[start : start+step]

				if got[:step] != expected {
					t.Logf("STEP=%d failed at batch %d (pos %d): got %s, want %s", step, batch, pos, got, expected)
					break
				}

				matchedBatches++
				pos += step
			}

			t.Logf("STEP=%d: successful batches = %d (accurate digits = %d)", step, matchedBatches, matchedBatches*step)

			if matchedBatches < 1 {
				t.Errorf("STEP=%d failed even first batch", step)
			}
		})
	}
}
