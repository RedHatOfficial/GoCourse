package main

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Headers check").Get("http://httpbin.org/get")
	f.Send()
	f.ExpectStatus(200)
	f.ExpectHeader("Server", "nginx")

	frisby.Global.PrintReport()
}
