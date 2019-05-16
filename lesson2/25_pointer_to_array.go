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

	p_third := &month[2]
	*p_third = "Březen"

	fmt.Println(month)
}
