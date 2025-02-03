package main

import (
	"fmt"

	"github.com/cosmos72/gomacro/fast"
)

const script = `
func multiply(a, b int) int {
    return a*b
}

x:=6
y:=7
multiply(x, y)
`

func RunGomacro(script string) any {
	interpreter := fast.New()
	vals, _ := interpreter.Eval(script)
	if len(vals) < 1 {
		return "no value"
	}
	return vals[0].ReflectValue()
}

func main() {
	fmt.Println(RunGomacro(script))
}
