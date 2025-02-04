package main

import (
	"fmt"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

const script = `
func multiply(a, b int) int {
    return a*b
}
`

func RunYaegi(script string) any {
	interpreter := interp.New(interp.Options{})

	interpreter.Use(stdlib.Symbols)

	ret, err := interpreter.Eval(script)
	if err != nil {
		panic(err)
	}

	return ret
}

func main() {
	fmt.Println(RunYaegi(script))
}
