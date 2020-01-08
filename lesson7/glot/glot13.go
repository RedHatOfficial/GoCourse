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

const points = 50

func main() {
	plot := NewPlot(2)
	defer plot.Close()

	var pointsX [points]float64
	for i := 0; i < points; i++ {
		pointsX[i] = float64(i) * 2.0 * math.Pi / points
	}

	function1 := func(t float64) float64 {
		// limita
		if t == 0.0 {
			return 1.0
		}
		return math.Sin(t) / t
	}

	plot.AddFunc2d("sinc t", "boxes", pointsX[:], function1)

	plot.SetTitle("Plot #13")
	plot.SetXLabel("t")
	plot.SetYLabel("y")

	plot.SetXrange(0, 1+int(math.Round(2.0*math.Pi)))
	plot.SetYrange(-1, 1)
	plot.CheckedCmd(`set xtics ('0' 0, '{/Symbol p}' pi, '2{/Symbol p}' 2*pi)`)
	plot.Save("glot13.png")
}
