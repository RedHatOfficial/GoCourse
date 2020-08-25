package main

import (
	"encoding/csv"
	"os"
)

// jméno generovaného souboru
const spreadsheetName = "test2.csv"

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
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

	// uložení všech řádků (záznamu) a provedení operace Flush
	writer.WriteAll(records)

	// test, zda při Write() nebo Flush() došlo k nějaké chybě
	err = writer.Error()
	if err != nil {
		panic(err)
	}
}
