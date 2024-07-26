package bbp

import (
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
	if start+num < CLIMIT && start >= 0 && num > 0 {
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

func (pi *PiDigits) GetDecimalValues(start int, num int) []byte {
	hex := pi.GetHexDigits(start, num)
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
