package main

import (
	"encoding/csv"
	"os"
)

// jméno generovaného souboru
const spreadsheetName = "test5.csv"

func main() {
	records := [][]string{
		{"first\nname", "last\nname", "username"},
		{"'Rob'", "'Pike'", "rob"},
		{"\"Ken\"", "\"Thompson\"", "ken"},
		{"`Robert`", "`Griesemer`", "gri"},
		{"A B", "C\tD", "\n"},
		{"Foo,Bar", "Baz,,,", ","},
	}

	// vytvoření výstupního souboru s tabulkou
	fout, err := os.Create(spreadsheetName)
	if err != nil {
		panic(err)
	}

	// zajištění, že se soubor s tabulkou uzavře při ukončení programu
	defer fout.Close()

	// konstrukce objektu pro postupné vytvoření tabulky
	writer := csv.NewWriter(fout)

	for _, record := range records {
		// vložení nového řádku (záznamu)
		err := writer.Write(record)
		if err != nil {
			panic(err)
		}
	}

	// pokus o uložení tabulky a vyprázdnění bufferu
	writer.Flush()

	// test, zda při Write() nebo Flush() došlo k nějaké chybě
	err = writer.Error()
	if err != nil {
		panic(err)
	}
}
