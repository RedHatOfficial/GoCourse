package main

import (
    "fmt"

    "github.com/k0kubun/pp/v3"
)

func __gore_p(xs ...any) {
    for _, x := range xs {
        pp.Println(x)
    }
}
func main() { _, _ = fmt.Println("Hello") }
