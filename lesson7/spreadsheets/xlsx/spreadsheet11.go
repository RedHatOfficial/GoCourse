package main

import (
	"fmt"

	"github.com/tealeg/xlsx/v3"
)

// jméno generovaného souboru
const spreadsheetName = "test11.xlsx"

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

	style := xlsx.NewStyle()
	style.Alignment.Horizontal = "right"
	style.Fill.FgColor = "ffa0ffa0"
	style.Fill.PatternType = "solid"
	style.Font.Bold = true
	style.ApplyAlignment = true
	style.ApplyFill = true
	style.ApplyFont = true

	// přidání první buňky na řádek a naplnění hodnotou
	cell := row.AddCell()
	cell.SetString("Test #1")

	// nastavení stylu buňky
	cell.SetStyle(style)

	style.Fill.FgColor = "ffffa0a0"
	style.Fill.PatternType = "solid"
	style.Font.Bold = false

	// přidání druhé buňky na řádek a naplnění hodnotou
	cell = row.AddCell()
	cell.SetString("Test #1")

	// nastavení stylu buňky
	cell.SetStyle(style)

	// pokus o uložení sešitu
	err = worksheet.Save(spreadsheetName)
	if err != nil {
		panic(err)
	}
}
