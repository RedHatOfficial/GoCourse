package main

import "fmt"

func computePi(n int, results chan<- float64) {
	pi := 4.0
	var i int
	for i = 3; i <= n+2; i += 2 {
		fi := float64(i)
		pi = pi * (fi - 1) / fi * (fi + 1) / fi
	}
	results <- pi
}

func main() {
	var results [30]chan float64
	for i := range results {
		results[i] = make(chan float64)
	}

	for i, n := 0, 1; n <= 100000000; i, n = i+1, n*2 {
		go computePi(n, results[i])
	}
	for i, n := 0, 1; n <= 100000000; i, n = i+1, n*2 {
		fmt.Printf("%d %16.14f\n", n, <-results[i])
	}
}
