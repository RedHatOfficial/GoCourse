package main

import "fmt"

func computePi(n int) float64 {
	pi := 4.0
	var i int
	for i = 3; i <= n+2; i += 2 {
		fi := float64(i)
		pi = pi * (fi - 1) / fi * (fi + 1) / fi
	}
	return pi
}

func main() {
	var n int
	for n = 1; n <= 1000000000; n *= 2 {
		fmt.Printf("%d %16.14f\n", n, computePi(n))
	}
}
