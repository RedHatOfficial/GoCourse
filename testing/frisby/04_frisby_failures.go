package main

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Simplest test").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)

	f = frisby.Create("Test that fails").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(404)

	f = frisby.Create("Another test that fails").Get("http://httpbin.org/status/456")
	f.Send()
	f.ExpectStatus(404)

	frisby.Global.PrintReport()
}
