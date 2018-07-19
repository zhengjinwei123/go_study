package mypath

func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

type K float32
type C float32

func KToC(k K) C {
	return C(-273.15 + k)
}

func CToK(c C) K {
	return K(c+273.15)
}
