package main

import (
	"fmt"
	"sort"
)

// Role represents a role in movie, drama etc.
type Role struct {
	name    string
	surname string
}

func printRoles(roles []Role) {
	for i, role := range roles {
		fmt.Printf("#%d: %s %s\n", i, role.name, role.surname)
	}
}

func main() {
	roles := []Role{
		{"Eliška", "Najbrtová"},
		{"Jenny", "Suk"},
		{"Anička", "Šafářová"},
		{"Sváťa", "Pulec"},
		{"Blažej", "Motyčka"},
		{"Eda", "Wasserfall"},
		{"Přemysl", "Hájek"},
	}

	fmt.Println("Unsorted:")
	printRoles(roles)
	fmt.Println("--------------------")

	sort.Slice(roles, func(i, j int) bool { return roles[i].name < roles[j].name })
	fmt.Println("Sorted by name:")
	printRoles(roles)
	fmt.Println("--------------------")

	sort.Slice(roles, func(i, j int) bool { return roles[i].surname < roles[j].surname })
	fmt.Println("Sorted by surname:")
	printRoles(roles)
	fmt.Println("--------------------")
}
