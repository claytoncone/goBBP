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

func (pi *PiDigits) GetHexValues(start, end int) []byte {
	hex := pi.GetHexDigits(start, end)
	chx := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}
	for i, v := range hex {
		hex[i] = chx[v]
	}
	return hex
}

func (pi *PiDigits) GetDecimalValues(start int, num int) []byte {
	hex := pi.GetHexDigits(start, num)
	out := make([]byte, len(hex))
	fraction := 0.0
	divisor := 1.0

	for i, h := range hex {
		divisor *= 1.6
		fraction += float64(int8(h)) / divisor

		digit := int(math.Floor(fraction))
		out[i] = byte(digit % 10) // Handle overflow with modulo

		fraction = (fraction - math.Floor(fraction)) * 10.0
	}

	return out
}
