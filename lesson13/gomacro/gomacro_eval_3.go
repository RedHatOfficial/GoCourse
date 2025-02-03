package main

import (
	"fmt"

	"github.com/cosmos72/gomacro/fast"
)

const script = `
x:=6
y:=7
z:=x*y
z
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
