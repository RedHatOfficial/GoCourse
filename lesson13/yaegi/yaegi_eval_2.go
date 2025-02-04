package main

import (
	"fmt"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

const script = `
x:=6
y:=7
z:=x*y
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
