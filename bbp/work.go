package bbp

import (
	"math"
)

type PiDigits struct {
	exponent *[NPT]float64
	channel  chan float64
}

func (pi *PiDigits) genHex(start int, num int) []byte {
	var d1 float64
	var d2 float64
	go pi.series(1, start, 4.)
	go pi.series(4, start, 2.)
	go pi.series(5, start, 1.)
	go pi.series(6, start, 1.)
	for i := 0; i < 4; i++ {
		d2 = <-pi.channel
		if d2 > 8. {

			d1 += d2
		} else {
			d1 -= d2
		}
	}
	pid := d1 - 10
	pid = pid - float64(int(pid)) + 1.
	return pi.ihex(pid, num)
}

func (pi *PiDigits) genExp() {
	var exp [NPT]float64
	exp[0] = 0.
	exp[1] = 1.
	for i := 2; i < NPT; i++ {
		exp[i] = 2. * exp[i-1]
	}
	pi.exponent = &exp
}

func (pi *PiDigits) series(m int, id int, kf float64) {
	var ak, p, s, t float64
	for k := 0; k < id; k++ {
		ak = 8*float64(k) + float64(m)
		p = float64(id) - float64(k)
		t = pi.expm(p, ak)
		s += t / ak
		s = s - float64(int(s))
	}
	for k := id; k <= id+100; k++ {
		ak = 8*float64(k) + float64(m)
		t = math.Pow(16., float64(id-k)) / ak

		if t < EPS {
			break
		}
		s += t
		s = s - float64(int(s))
	}

	s = s * kf

	if kf == 4. {
		s += 10.
	}

	pi.channel <- s
}

func (pi *PiDigits) expm(p float64, ak float64) float64 {
	var p1, pt, r float64

	if ak == 1. {
		return 0.
	}

	i := 0
	for ; i < NPT; i++ {
		if pi.exponent[i] > p {
			break
		}
	}
	pt = pi.exponent[i-1]
	p1 = p
	r = 1
	for j := 1; j < i; j++ {
		if p1 >= pt {
			r = 16. * r
			r = r - float64(float64(int(r/ak))*ak)
			p1 = p1 - pt
		}

		pt = 0.5 * pt

		if pt >= 1. {
			r = r * r
			r = r - float64(float64(int(r/ak))*ak)
		}
	}

	return r
}

func (pi *PiDigits) ihex(x float64, num int) []byte {
	var out []byte
	var y float64
	y = math.Abs(x)

	for i := 0; i < num; i++ {
		y = (y - math.Floor(y)) * 16
		out = append(out, byte(y))
	}

	return out
}