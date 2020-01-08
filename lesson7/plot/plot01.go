package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
)

const resX = 20.0 / 3.0 * vg.Inch
const resY = 5.0 * vg.Inch

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	err = p.Save(resX, resY, "plot01.png")
	if err != nil {
		panic(err)
	}
}
