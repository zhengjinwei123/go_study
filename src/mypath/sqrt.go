package mypath


func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

func Add(a int,b int)(int,int){
	return a+b,a-b
}





