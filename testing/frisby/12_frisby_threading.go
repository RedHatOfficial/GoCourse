package main

import "github.com/verdverm/frisby"

func main() {
	frisby.Create("Cookies check").
		Get("http://httpbin.org/cookies").
		SetCookie("foo", "bar").
		Send().
		ExpectStatus(200).
		ExpectJson("cookies.foo", "bar").
		PrintBody()

	frisby.Global.PrintReport()
}
