package main

import "github.com/k0kubun/pp/v3"

func __gore_p(xs ...any) {
    for _, x := range xs {
        pp.Println(x)
    }
}
func main() { x := 6; y := 7; z := "foobar" }
