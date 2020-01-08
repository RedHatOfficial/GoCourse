package main

import (
	"github.com/Arafatk/glot"
	"math"
)

type Plot struct {
	*glot.Plot
}

func (plot *Plot) Save(filename string) {
	plot.Cmd("set terminal pngcairo")
	plot.Cmd("set output '" + filename + "'")
	plot.Cmd("replot")
}

func NewPlot(dimensions int) *Plot {
	plot, err := glot.NewPlot(dimensions, false, false)
	if err != nil {
		panic(err)
	}
	return &Plot{plot}

}

const points = 100

func main() {
	plot := NewPlot(2)
	defer plot.Close()

	var pointsX [points]float64
	for i := 0; i < points; i++ {
		pointsX[i] = float64(i) * 2.0 * math.Pi / points
	}

	function1 := func(t float64) float64 {
		return math.Sin(t)
	}

	function2 := func(t float64) float64 {
		return math.Sin(t + 2.0*math.Pi/3)
	}

	function3 := func(t float64) float64 {
		return math.Sin(t - 2.0*math.Pi/3)
	}

	plot.AddFunc2d("sin t", "lines", pointsX[:], function1)
	plot.AddFunc2d("sin t+2{/Symbol p}/3", "lines", pointsX[:], function2)
	plot.AddFunc2d("sin t-2{/Symbol p}/3", "lines", pointsX[:], function3)

	plot.SetTitle("Plot #12")
	plot.SetXLabel("t")
	plot.SetYLabel("y")

	plot.SetXrange(0, 1+int(math.Round(2.0*math.Pi)))
	plot.SetYrange(-1, 1)
	plot.CheckedCmd(`set xtics ('0' 0, '{/Symbol p}' pi, '2{/Symbol p}' 2*pi)`)
	plot.Save("glot12.png")
}
