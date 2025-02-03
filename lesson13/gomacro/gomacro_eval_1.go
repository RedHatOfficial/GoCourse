package main

import (
	"fmt"

	"github.com/cosmos72/gomacro/fast"
)

func RunGomacro(script string) any {
	interpreter := fast.New()
	vals, _ := interpreter.Eval(script)
	return vals[0].ReflectValue()
}

func main() {
	fmt.Println(RunGomacro("6*7"))
}
