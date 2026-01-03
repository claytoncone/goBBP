package bbp

import (
	"goBBP/internal/convert"
	"math"
	"runtime"
)

func New() *PiDigits {
	pi := &PiDigits{}
	pi.channel = make(chan float64)
	return pi
}

func (pi *PiDigits) GetHexDigits(start int, num int) []byte {
	var out []byte
	if start <= CLIMIT && start+num < CLIMIT && start >= 0 && num > 0 {
		numCpu := runtime.NumCPU()
		runtime.GOMAXPROCS(numCpu)

		for i := 0; i < num; i = i + STEP {
			c := num - i
			if c > STEP {
				out = append(out, pi.genHex(start+i, STEP)...)
			} else {
				out = append(out, pi.genHex(start+i, c)...)
			}
		}
	}
	return out
}

func (pi *PiDigits) GetHexValues(start, num int) []byte {
	hex := pi.GetHexDigits(start, num)
	return convert.HexBytesToAlpha(hex)
}

// GetHexString returns hex digits as an uppercase string
func (pi *PiDigits) GetHexString(start int, num int) string {
	hex := pi.GetHexDigits(start, num)
	return convert.HexBytesToString(hex)
}

// GetHexLowerString returns hex digits as a lowercase string
func (pi *PiDigits) GetHexLowerString(start int, num int) string {
	hex := pi.GetHexDigits(start, num)
	return convert.HexBytesToLowerString(hex)
}

// GetDecimalValues converts hex digits to decimal digits.
//
// LIMITATION: Due to float64 precision, only the first ~18 decimal digits
// are reliable. This is suitable for display purposes but not for
// generating large quantities of decimal digits.
//
// TODO: Implement proper hex-fraction to decimal-fraction conversion
// that can handle arbitrary lengths (similar to series() reset approach).
func (pi *PiDigits) GetDecimalValues(start int, num int) []byte {
	hex := pi.GetHexDigits(start, num)
	if len(hex) == 0 {
		return nil
	}

	out := make([]byte, len(hex))
	fraction := 0.0
	divisor := 1.0

	for i, h := range hex {
		divisor *= 1.6
		fraction += float64(int8(h)) / divisor
		out[i] = byte(fraction)
		fraction = (fraction - math.Floor(fraction)) * 10.0
	}

	return out
}

// GetDecimalString returns decimal digits as a string.
// See GetDecimalValues for precision limitations.
func (pi *PiDigits) GetDecimalString(start int, num int) string {
	dec := pi.GetDecimalValues(start, num)
	return convert.DecBytesToString(dec)
}
