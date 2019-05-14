package main

import (
	"fmt"
	"sort"
)

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
		Role{"Eliška", "Najbrtová"},
		Role{"Jenny", "Suk"},
		Role{"Anička", "Šafářová"},
		Role{"Sváťa", "Pulec"},
		Role{"Blažej", "Motyčka"},
		Role{"Eda", "Wasserfall"},
		Role{"Přemysl", "Hájek"},
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
