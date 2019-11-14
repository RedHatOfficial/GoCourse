package accumulator

import (
	"flag"
	"github.com/DATA-DOG/godog"
)

type acc struct {
	value int
}

func (a *acc) accumulate(x int) {
	a.value += x
}

var opt = godog.Options{
	Format: "progress",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}
