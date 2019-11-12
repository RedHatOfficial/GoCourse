package main

import "github.com/verdverm/frisby"

func main() {
	f := frisby.Create("Content check").Get("http://golang.org")
	f.Send()
	f.ExpectStatus(200)
	f.ExpectHeader("Content-Type", "text/html; charset=utf-8").ExpectContent("The Go Programming Language")

	data := map[string]string{
		"text": "Hello **world**!",
	}
	f = frisby.Create("Markdown conversion").Post("https://api.github.com/markdown").SetJson(data)
	f.Send()
	f.ExpectStatus(200)
	f.ExpectContent("<p>Hello <strong>world</strong>!</p>")

	f.PrintBody()

	frisby.Global.PrintReport()
}
