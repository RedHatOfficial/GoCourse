package main

import (
	"fmt"

	"github.com/cosmos72/gomacro/fast"
)

const script = `
x:=6
y:=7
z:=x*y
`

func RunGomacro(script string) any {
	interp := fast.New()
	vals, _ := interp.Eval(script)
	if len(vals) < 1 {
		return "no value"
	}
	return vals[0].ReflectValue()
}

func main() {
	fmt.Println(RunGomacro(script))
}
