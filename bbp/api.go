package bbp

import (
	"runtime"
)

func New() *PiDigits {
	pi := &PiDigits{}
	pi.genExp()
	pi.channel = make(chan float64)

	return pi
}

func (pi *PiDigits) Get(start int, num int) []byte {
	var out []byte
	if start <= CLIMIT && start+num < CLIMIT && start >= 0 && num > 0 {
		numcpu := runtime.NumCPU()
		runtime.GOMAXPROCS(numcpu)

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
