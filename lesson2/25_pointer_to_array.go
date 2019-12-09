package main

import "fmt"

func main() {
	month := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December"}

	fmt.Println(month)

	// cont

	pThird := &month[2]
	*pThird = "Březen"

	fmt.Println(month)
}
