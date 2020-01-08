package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"math/rand"
)

const resX = 20.0 / 3.0 * vg.Inch
const resY = 5.0 * vg.Inch

func main() {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	input := [...]int32{1, 2, 4, 8, 9, 8, 4, 2, 1}

	points1 := make(plotter.XYs, len(input))
	points2 := make(plotter.XYs, len(input))
	for i := range input {
		points1[i].X = float64(i)
		points2[i].X = float64(i)
		points1[i].Y = float64(input[i])
		points2[i].Y = float64(input[i]) + rand.Float64()/2.0 - 0.5
	}

	p.Title.Text = "Plot #3"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p, "Expected data", points1, "Measured data", points2)
	if err != nil {
		panic(err)
	}

	err = p.Save(resX, resY, "plot04.png")
	if err != nil {
		panic(err)
	}
}
