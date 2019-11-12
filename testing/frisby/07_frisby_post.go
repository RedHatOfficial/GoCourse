package main

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Simplest test").Post("http://httpbin.org/post")
	f.Send()
	f.ExpectStatus(200)

	frisby.Global.PrintReport()
}
