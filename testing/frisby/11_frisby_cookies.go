package main

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Cookies check").Get("http://httpbin.org/cookies").SetCookie("foo", "bar")
	f.Send()
	f.ExpectStatus(200)
	f.ExpectJson("cookies.foo", "bar")

	f.PrintBody()

	frisby.Global.PrintReport()
}
