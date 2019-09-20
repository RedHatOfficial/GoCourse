package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["first"] = "1st"
	m["second"] = "2nd"
	m["third"] = "3rd"
	fmt.Printf("'%s'\n", m["first"])
	fmt.Printf("'%s'\n", m["second"])
	fmt.Printf("'%s'\n", m["third"])
	fmt.Printf("'%s'\n", m["xyzzy"])
}
