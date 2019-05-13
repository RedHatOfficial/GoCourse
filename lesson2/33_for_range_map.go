package main

func main() {
	var m1 map[int]string = make(map[int]string)
	m1[0] = "one"
	m1[1] = "two"
	m1[2] = "three"
	m1[3] = "four"

	for key, val := range m1 {
		println(key, val)
	}
}
