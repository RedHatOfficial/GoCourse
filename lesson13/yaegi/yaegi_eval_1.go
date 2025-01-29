package main

import (
	"fmt"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

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
	fmt.Println(RunYaegi("6*7"))
}
