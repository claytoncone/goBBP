package bbp

import (
	"math"
)

type PiDigits struct {
	channel chan float64
}

func (pi *PiDigits) genHex(start int, num int) []byte {
	var pid float64
	go pi.series(1, start, 4.)
	go pi.series(4, start, -2.)
	go pi.series(5, start, -1.)
	go pi.series(6, start, -1.)
	for i := 0; i < 4; i++ {
		pid += <-pi.channel
	}
	pid -= math.Floor(pid)
	return pi.ihex(pid, num)
}

func (pi *PiDigits) series(m int, id int, kf float64) {
	var s, t float64

	for k := 0; k < id; k++ {
		ak := float64(8*k + m)
		p := float64(id - k)
		t = pi.expm(p, ak)
		s += t / ak
		s -= math.Floor(s)
	}

	for k := id; k <= id+100; k++ {
		ak := float64(8*k + m)
		t := math.Pow(16., float64(id-k)) / ak

		if t < EPS {
			break
		}

		s += t
		s -= math.Floor(s)
	}
	s *= kf

	pi.channel <- s
}

func (pi *PiDigits) expm(p float64, ak float64) float64 {
	var p1, pt, r float64

	if ak == 1. {
		return 0.
	}

	i := 0
	for ; i < NPT; i++ {
		if Exponents[i] > p {
			break
		}
	}
	pt = Exponents[i-1]
	p1 = p
	r = 1.
	for j := 1; j <= i; j++ {
		if p1 >= pt {
			r *= 16.
			r -= float64(float64(int(r/ak)) * ak)
			p1 -= pt
		}

		pt /= 2.

		if pt >= 1. {
			r *= r
			r -= float64(float64(int(r/ak)) * ak)
		}
	}

	return r
}

func (pi *PiDigits) ihex(x float64, num int) []byte {
	var out []byte
	var y float64
	y = math.Abs(x + 1.)

	for i := 0; i < num; i++ {
		y = (y - math.Floor(y)) * 16
		out = append(out, byte(y))
	}

	return out
}
