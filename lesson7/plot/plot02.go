package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

const resX = 20.0 / 3.0 * vg.Inch
const resY = 5.0 * vg.Inch

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	input := [...]int32{1, 2, 4, 8, 9, 8, 4, 2, 1}

	points := make(plotter.XYs, len(input))
	for i := range points {
		points[i].X = float64(i)
		points[i].Y = float64(input[i])
	}
	err = plotutil.AddLinePoints(p, "Measured data", points)
	if err != nil {
		panic(err)
	}

	err = p.Save(resX, resY, "plot02.png")
	if err != nil {
		panic(err)
	}
}
