package bbp

const (
	STEP = 8
)

const (
	CLIMIT = 10000000
	NPT    = 25
	EPS    = 1e-17
)

var Exponents = []float64{ // populated with successive powers of 2
	0., 1., 2., 4., 8., 16., 32., 64., 128., 256.,
	512., 1024., 2048., 4096., 8192., 16384., 32768., 65536., 131072., 262144.,
	524288., 1048576., 2097152., 4194304., 8388608.}
