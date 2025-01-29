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
	i := interp.New(interp.Options{})

	i.Use(stdlib.Symbols)

	ret, err := i.Eval(script)
	if err != nil {
		panic(err)
	}

	return ret
}

func main() {
	fmt.Println(RunYaegi(script))
}
