package main

// #cgo LDFLAGS: -lm
//#include <math.h>
import "C"
import "fmt"
import "math"

var data = []float64{
	1.0,
	0.0,
	0.9220181640,
	0.6975023699,
	0.9505543193,
	0.1101513685,
	0.6499663974,
	0.4060567980,
	0.0025370125,
	0.1790128792,
}

// You can't use CGo in tests
func Csin(num float64) float64 {
	return float64(C.sin(C.double(num)))
}

func Ccos(num float64) float64 {
	return float64(C.cos(C.double(num)))
}

// Ctgamma is a proxy to C tgamma function
func Ctgamma(num float64) float64 {
	return float64(C.tgamma(C.double(num)))
}

func main() {
	fmt.Println("Go math")
	for i := 0; i < len(data); i++ {
		fmt.Printf("Number: %f Sin: %f Cos: %f Tan: %f\n", data[i], math.Sin(data[i]), math.Cos(data[i]), math.Tan(data[i]))
		fmt.Printf("Asin: %f ACos: %f ATan: %f Log: %f \n", math.Asin(data[i]), math.Acos(data[i]), math.Atan(data[i]), math.Log(data[i]))
		fmt.Printf("Log10: %f Sinh: %f Cosh: %f Tanh: %f\n", math.Log10(data[i]), math.Sinh(data[i]), math.Cosh(data[i]), math.Tanh(data[i]))
		fmt.Printf("Erf: %f tGamma: %f\n", math.Erf(data[i]), math.Gamma(data[i]))
	}
	fmt.Println("libc math")
	for i := 0; i < len(data); i++ {
		num := C.double(data[i])
		fmt.Printf("Number: %f Sin: %f Cos: %f Tan: %f\n", data[i], Csin(data[i]), Ccos(data[i]), C.tan(num))
		fmt.Printf("Asin: %f ACos: %f ATan: %f Log: %f \n", C.asin(num), C.acos(num), C.atan(num), C.log(num))
		fmt.Printf("Log10: %f Sinh: %f Cosh: %f Tanh: %f\n", C.log10(num), C.sinh(num), C.cosh(num), C.tanh(num))
		fmt.Printf("Erf: %f tGamma: %f\n", C.erf(num), Ctgamma(data[i]))
	}
}
