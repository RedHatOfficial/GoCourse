package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test06.xlsx"

func main() {
	// konstrukce struktury typu File
	worksheet := xlsx.NewFile()

	fmt.Println(worksheet)

	// přidání listu do sešitu
	sheet, err := worksheet.AddSheet("Tabulka1")
	if err != nil {
		panic(err)
	}
	defer sheet.Close()

	fmt.Println(sheet)

	// přidání řádku do tabulky
	row := sheet.AddRow()

	// přidání buňky na řádek a naplnění hodnotou
	row.AddCell().SetString("Hello")
	row.AddCell().SetInt(42)
	row.AddCell().SetInt64(1000 * 1000 * 1000 * 1000)
	row.AddCell().SetFloat(1 / 3.0)
	row.AddCell().SetBool(true)
	row.AddCell().SetBool(false)
	row.AddCell().SetDate(time.Now())
	row.AddCell().SetDateTime(time.Now())

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
