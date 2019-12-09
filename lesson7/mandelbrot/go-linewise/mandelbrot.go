package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func writeImage(width uint, height uint, image []byte) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintln(w, "P3")
	fmt.Fprintf(w, "%d %d\n", width, height)
	fmt.Fprintln(w, "255")

	for i := 0; i < len(image); {
		r := image[i]
		i++
		g := image[i]
		i++
		b := image[i]
		i++
		fmt.Fprintf(w, "%d %d %d\n", r, g, b)
	}
}

func calcMandelbrot(width uint, height uint, maxiter uint, palette [][3]byte, image []byte, cy float64) {
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
		image[3*x] = color[0]
		image[3*x+1] = color[1]
		image[3*x+2] = color[2]
		cx += 3.0 / float64(width)
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

	image := make([]byte, width*height*3)
	offset := 0
	delta := width * 3

	var cy float64 = -1.5
	for y := 0; y < height; y++ {
		calcMandelbrot(uint(width), uint(height), uint(maxiter), mandmap[:], image[offset:offset+delta], cy)
		offset += delta
		cy += 3.0 / float64(height)
	}
	writeImage(uint(width), uint(height), image)
}
