package main

import (
	"fmt"
	"os"
	"strconv"
)

func calcMandelbrot(width uint, height uint, maxiter uint, palette [][3]byte) {
	fmt.Println("P3")
	fmt.Printf("%d %d\n", width, height)
	fmt.Println("255")

	var cy float64 = -1.5

	for y := uint(0); y < height; y++ {
		var cx float64 = -2.0
		for x := uint(0); x < width; x++ {
			var zx float64 = 0.0
			var zy float64 = 0.0
			var i uint = 0
			for i < maxiter {
				zx2 := zx * zx
				zy2 := zy * zy
				if zx2+zy2 > 4.0 {
					break
				}
				zy = 2.0*zx*zy + cy
				zx = zx2 - zy2 + cx
				i++
			}
			color := palette[i]
			r := color[0]
			g := color[1]
			b := color[2]
			fmt.Printf("%d %d %d\n", r, g, b)
			cx += 3.0 / float64(width)
		}
		cy += 3.0 / float64(height)
	}

}

func main() {
	if len(os.Args) < 4 {
		println("usage: ./mandelbrot width height maxiter")
		os.Exit(1)
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Improper width parameter: '%s'\n", os.Args[1])
		os.Exit(1)
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Improper height parameter: '%s'\n", os.Args[2])
		os.Exit(1)
	}

	maxiter, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("Improper maxiter parameter: '%s'\n", os.Args[3])
		os.Exit(1)
	}
	calcMandelbrot(uint(width), uint(height), uint(maxiter), mandmap[:])
}
