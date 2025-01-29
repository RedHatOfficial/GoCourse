package main

import (
	"fmt"

	"github.com/cosmos72/gomacro/fast"
)

func RunGomacro(script string) any {
	interp := fast.New()
	vals, _ := interp.Eval(script)
	return vals[0].ReflectValue()
}

func main() {
	fmt.Println(RunGomacro("6*7"))
}
